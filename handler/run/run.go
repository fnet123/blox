// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package run

import (
	"context"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	"github.com/pkg/errors"

	"github.com/aws/amazon-ecs-event-stream-handler/handler/api/v1"
	"github.com/aws/amazon-ecs-event-stream-handler/handler/clients"
	"github.com/aws/amazon-ecs-event-stream-handler/handler/event"
	"github.com/aws/amazon-ecs-event-stream-handler/handler/reconcile"
	"github.com/aws/amazon-ecs-event-stream-handler/handler/store"
	"github.com/urfave/negroni"
)

const (
	serverAddress     = "localhost:3000"
	serverReadTimeout = 10 * time.Second
)

// StartEventStreamHandler starts the event stream handler. It creates an ETCD
// client, a data store using this client and an event processor to process
// events from the SQS queue. It also starts the RESTful server and blocks on
// the listen method of the same to listen to requests that query for task and
// instance state from the store.
func StartEventStreamHandler(queueName string, etcdEndpoints []string) error {
	etcdClient, err := clients.NewEtcdClient(etcdEndpoints)
	if err != nil {
		return errors.Wrapf(err, "Could not start etcd")
	}
	defer etcdClient.Close()

	// initialize the datastore
	datastore, err := store.NewDataStore(etcdClient)
	if err != nil {
		return errors.Wrapf(err, "Could not initialize the datastore")
	}

	// initialize services
	stores, err := store.NewStores(datastore)
	if err != nil {
		return errors.Wrapf(err, "Could not initialize stores")
	}

	awsSession, err := clients.NewAWSSession()
	if err != nil {
		return errors.Wrapf(err, "Could not load aws session")
	}

	ecsClient := clients.NewECSClient(awsSession)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	recon, err := reconcile.NewReconciler(ctx, stores, ecsClient, reconcile.ReconcileDuration)
	if err != nil {
		return errors.Wrapf(err, "Could not start reconciler")
	}
	err = recon.RunOnce()
	if err != nil {
		return errors.Wrapf(err, "Error bootstrapping")
	}
	log.Infof("Bootstrapping completed")
	go recon.Run()

	// initialize apis
	apis := v1.NewAPIs(stores)

	// start event processor
	processor := event.NewProcessor(stores)

	sqsClient := clients.NewSQSClient(awsSession)

	// start event consumer
	consumer, err := event.NewConsumer(sqsClient, processor, queueName)
	if err != nil {
		return errors.Wrapf(err, "Could not start the consumer")
	}

	go consumer.PollForEvents(ctx)

	// start server
	router := v1.NewRouter(apis)

	n := negroni.Classic()
	n.UseHandler(router)

	s := &http.Server{
		Addr:        serverAddress,
		Handler:     n,
		ReadTimeout: serverReadTimeout,
	}

	return s.ListenAndServe()
}
