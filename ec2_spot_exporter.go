// Copyright 2020 Patrick Jonathan Cadelina
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Namespace   = "ec2_spot"
	MetricsPath = "/metrics"
)

type Collector struct {
	spotInstanceRequests                                  *prometheus.Desc
	spotInstanceAzGroupConstraintRequests                 *prometheus.Desc
	spotInstanceBadParametersRequests                     *prometheus.Desc
	spotInstanceCanceledBeforeFulfillmentRequests         *prometheus.Desc
	spotInstanceCapacityNotAvailableRequests              *prometheus.Desc
	spotInstanceConstraintNotFulfillableRequests          *prometheus.Desc
	spotInstanceFulfilledRequests                         *prometheus.Desc
	spotInstanceStoppedByPriceRequests                    *prometheus.Desc
	spotInstanceStoppedByUserRequests                     *prometheus.Desc
	spotInstanceStoppedNoCapacityRequests                 *prometheus.Desc
	spotInstanceTerminatedByPriceRequests                 *prometheus.Desc
	spotInstanceTerminatedByScheduleRequests              *prometheus.Desc
	spotInstanceTerminatedByServiceRequests               *prometheus.Desc
	spotInstanceTerminatedByUserRequests                  *prometheus.Desc
	spotInstanceTerminatedLaunchGroupConstraintRequests   *prometheus.Desc
	spotInstanceTerminatedNoCapacityRequests              *prometheus.Desc
	spotInstanceLaunchGroupConstraintRequests             *prometheus.Desc
	spotInstanceMarkedForStopRequests                     *prometheus.Desc
	spotInstanceMarkedForTerminationRequests              *prometheus.Desc
	spotInstanceNotScheduledYetRequests                   *prometheus.Desc
	spotInstancePendingEvaluationRequests                 *prometheus.Desc
	spotInstancePendingFulfillmentRequests                *prometheus.Desc
	spotInstancePlacementGroupConstraintRequests          *prometheus.Desc
	spotInstancePriceTooLowRequests                       *prometheus.Desc
	spotInstanceRequestCanceledAndInstanceRunningRequests *prometheus.Desc
	spotInstanceScheduleExpiredRequests                   *prometheus.Desc
	spotInstanceSystemErrorRequests                       *prometheus.Desc
}

func NewCollector() *Collector {
	return &Collector{
		spotInstanceRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_requests"),
			"Spot instance requests count.",
			[]string{"region"},
			nil,
		),
		spotInstanceAzGroupConstraintRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_az_group_contraint_requests"),
			"Spot instance requests with az-group-constraint status count",
			[]string{"region"},
			nil,
		),
		spotInstanceBadParametersRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_bad_parameters_requests"),
			"Spot instance requests with az-group-constraint status count",
			[]string{"region"},
			nil,
		),
		spotInstanceCanceledBeforeFulfillmentRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_canceled_before_fulfillment_requests"),
			"Spot instance requests with canceled-before-fulfillment status count",
			[]string{"region"},
			nil,
		),
		spotInstanceCapacityNotAvailableRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_capacity_not_available_requests"),
			"Spot instance requests with capacity-not-available status count",
			[]string{"region"},
			nil,
		),
		spotInstanceConstraintNotFulfillableRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_constraint_not_fulfillable_requests"),
			"Spot instance requests with constraint-not-fulfillable status count",
			[]string{"region"},
			nil,
		),
		spotInstanceFulfilledRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_fulfilled_requests"),
			"Spot instance requests with fulfilled status count",
			[]string{"region"},
			nil,
		),
		spotInstanceStoppedByPriceRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_stopped_by_price_requests"),
			"Spot instance requests with instance-stopped-by-price status count",
			[]string{"region"},
			nil,
		),
		spotInstanceStoppedByUserRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_stopped_by_user_requests"),
			"Spot instance requests with instance-stopped-by-user status count",
			[]string{"region"},
			nil,
		),
		spotInstanceStoppedNoCapacityRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_stopped_no_capacity_requests"),
			"Spot instance requests with instance-stopped-no-capacity status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedByPriceRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_by_price_requests"),
			"Spot instance requests with instance-terminated-by-price status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedByScheduleRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_by_schedule_requests"),
			"Spot instance requests with instance-terminated-by-schedule status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedByServiceRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_by_service_requests"),
			"Spot instance requests with instance-terminated-by-service status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedByUserRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_by_user_requests"),
			"Spot instance requests with instance-terminated-by-user status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedLaunchGroupConstraintRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_launch_group_constraint_requests"),
			"Spot instance requests with instance-terminated-launch-group-constraint status count",
			[]string{"region"},
			nil,
		),
		spotInstanceTerminatedNoCapacityRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_terminated_no_capacity_requests"),
			"Spot instance requests with instance-terminated-no-capacity status count",
			[]string{"region"},
			nil,
		),
		spotInstanceLaunchGroupConstraintRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_launch_group_constraint_requests"),
			"Spot instance requests with launch-group-constraint status count",
			[]string{"region"},
			nil,
		),
		spotInstanceMarkedForStopRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_marked_for_stop_requests"),
			"Spot instance requests with marked-for-stop status count",
			[]string{"region"},
			nil,
		),
		spotInstanceMarkedForTerminationRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_marked_for_termination_requests"),
			"Spot instance requests with marked-for-termination status count",
			[]string{"region"},
			nil,
		),
		spotInstanceNotScheduledYetRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_not_scheduled_yet_requests"),
			"Spot instance requests with not-scheduled-yet status count",
			[]string{"region"},
			nil,
		),
		spotInstancePendingEvaluationRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_pending_evaluation_requests"),
			"Spot instance requests with pending-evaluation status count.",
			[]string{"region"},
			nil,
		),
		spotInstancePendingFulfillmentRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_pending_fulfillment_requests"),
			"Spot instance requests with pending-fulfillment status count",
			[]string{"region"},
			nil,
		),
		spotInstancePlacementGroupConstraintRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_placement_group_constraint_requests"),
			"Spot instance requests with placement-group-constraint status count",
			[]string{"region"},
			nil,
		),
		spotInstancePriceTooLowRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_price_too_low_requests"),
			"Spot instance requests with price-too-low status count",
			[]string{"region"},
			nil,
		),
		spotInstanceRequestCanceledAndInstanceRunningRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_request_canceled_and_instance_running_requests"),
			"Spot instance requests with request-canceled-and-instance-running status count",
			[]string{"region"},
			nil,
		),
		spotInstanceScheduleExpiredRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_schedule_expired_requests"),
			"Spot instance requests with schedule_expired status count",
			[]string{"region"},
			nil,
		),
		spotInstanceSystemErrorRequests: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "instance_system_error_requests"),
			"Spot instance requests with system-error status count",
			[]string{"region"},
			nil,
		),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.spotInstanceRequests
	ch <- c.spotInstanceAzGroupConstraintRequests
	ch <- c.spotInstanceBadParametersRequests
	ch <- c.spotInstanceCanceledBeforeFulfillmentRequests
	ch <- c.spotInstanceCapacityNotAvailableRequests
	ch <- c.spotInstanceConstraintNotFulfillableRequests
	ch <- c.spotInstanceFulfilledRequests
	ch <- c.spotInstanceStoppedByPriceRequests
	ch <- c.spotInstanceStoppedByUserRequests
	ch <- c.spotInstanceStoppedNoCapacityRequests
	ch <- c.spotInstanceTerminatedByPriceRequests
	ch <- c.spotInstanceTerminatedByScheduleRequests
	ch <- c.spotInstanceTerminatedByServiceRequests
	ch <- c.spotInstanceTerminatedByUserRequests
	ch <- c.spotInstanceTerminatedLaunchGroupConstraintRequests
	ch <- c.spotInstanceTerminatedNoCapacityRequests
	ch <- c.spotInstanceLaunchGroupConstraintRequests
	ch <- c.spotInstanceMarkedForStopRequests
	ch <- c.spotInstanceMarkedForTerminationRequests
	ch <- c.spotInstanceNotScheduledYetRequests
	ch <- c.spotInstancePendingEvaluationRequests
	ch <- c.spotInstancePendingFulfillmentRequests
	ch <- c.spotInstancePlacementGroupConstraintRequests
	ch <- c.spotInstancePriceTooLowRequests
	ch <- c.spotInstanceRequestCanceledAndInstanceRunningRequests
	ch <- c.spotInstanceScheduleExpiredRequests
	ch <- c.spotInstanceSystemErrorRequests
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	sess := session.Must(session.NewSession())

	svc := ec2.New(sess)
	input := &ec2.DescribeSpotInstanceRequestsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("state"),
				Values: []*string{
					aws.String("open"),
					aws.String("active"),
					aws.String("closed"),
					aws.String("cancelled"),
					aws.String("failed"),
				},
			},
		},
	}

	result, err := svc.DescribeSpotInstanceRequests(input)
	if err != nil {
		log.Fatal(err)
	}

	var azGroupConstraintRequests float64 = 0
	var badParametersRequests float64 = 0
	var canceledBeforeFulfillmentRequests float64 = 0
	var capacityNotAvailableRequests float64 = 0
	var constraintNotFulfillableRequests float64 = 0
	var fulfilledRequests float64 = 0
	var instanceStoppedByPriceRequests float64 = 0
	var instanceStoppedByUserRequests float64 = 0
	var instanceStoppedNoCapacityRequests float64 = 0
	var instanceTerminatedByPriceRequests float64 = 0
	var instanceTerminatedByScheduleRequests float64 = 0
	var instanceTerminatedByServiceRequests float64 = 0
	var instanceTerminatedByUserRequests float64 = 0
	var instanceTerminatedLaunchGroupConstraintRequests float64 = 0
	var instanceTerminatedNoCapacityRequests float64 = 0
	var launchGroupConstraintRequests float64 = 0
	var markedForStopRequests float64 = 0
	var markedForTerminationRequests float64 = 0
	var notScheduledYetRequests float64 = 0
	var pendingEvaluationRequests float64 = 0
	var pendingFulfillmentRequests float64 = 0
	var placementGroupConstraintRequests float64 = 0
	var priceTooLowRequests float64 = 0
	var requestCanceledAndInstanceRunningRequests float64 = 0
	var scheduleExpiredRequests float64 = 0
	var systemErrorRequests float64 = 0

	for i := 0; i < len(result.SpotInstanceRequests); i++ {
		switch *result.SpotInstanceRequests[i].Status.Code {
		case "az-group-constraint":
			azGroupConstraintRequests += 1
		case "bad-parameters":
			badParametersRequests += 1
		case "canceled-before-fulfillment":
			canceledBeforeFulfillmentRequests += 1
		case "capacity-not-available":
			capacityNotAvailableRequests += 1
		case "constraint-not-fulfillable":
			constraintNotFulfillableRequests += 1
		case "fulfilled":
			fulfilledRequests += 1
		case "instance-stopped-by-price":
			instanceStoppedByPriceRequests += 1
		case "instance-stopped-by-user":
			instanceStoppedByUserRequests += 1
		case "instance-stopped-no-capacity":
			instanceStoppedNoCapacityRequests += 1
		case "instance-terminated-by-price":
			instanceTerminatedByPriceRequests += 1
		case "instance-terminated-by-schedule":
			instanceTerminatedByScheduleRequests += 1
		case "instance-terminated-by-service":
			instanceTerminatedByServiceRequests += 1
		case "instance-terminated-by-user":
			instanceTerminatedByUserRequests += 1
		case "instance-terminated-launch-group-constraint":
			instanceTerminatedLaunchGroupConstraintRequests += 1
		case "instance-terminated-no-capacity":
			instanceTerminatedNoCapacityRequests += 1
		case "launch-group-constraint":
			launchGroupConstraintRequests += 1
		case "marked-for-stop":
			markedForStopRequests += 1
		case "marked-for-termination":
			markedForTerminationRequests += 1
		case "not-scheduled-yet":
			notScheduledYetRequests += 1
		case "pending-evaluation":
			pendingEvaluationRequests += 1
		case "pending-fulfillment":
			pendingFulfillmentRequests += 1
		case "placement-group-constraint":
			placementGroupConstraintRequests += 1
		case "price-too-low":
			priceTooLowRequests += 1
		case "request-canceled-and-instance-running":
			requestCanceledAndInstanceRunningRequests += 1
		case "schedule-expired":
			scheduleExpiredRequests += 1
		case "system-error":
			systemErrorRequests += 1
		}
	}

	region := os.Getenv("AWS_REGION")

	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceRequests,
		prometheus.GaugeValue,
		float64(len(result.SpotInstanceRequests)),
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceAzGroupConstraintRequests,
		prometheus.GaugeValue,
		azGroupConstraintRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceBadParametersRequests,
		prometheus.GaugeValue,
		badParametersRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceCanceledBeforeFulfillmentRequests,
		prometheus.GaugeValue,
		canceledBeforeFulfillmentRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceCapacityNotAvailableRequests,
		prometheus.GaugeValue,
		capacityNotAvailableRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceConstraintNotFulfillableRequests,
		prometheus.GaugeValue,
		constraintNotFulfillableRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceFulfilledRequests,
		prometheus.GaugeValue,
		fulfilledRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceStoppedByPriceRequests,
		prometheus.GaugeValue,
		instanceStoppedByPriceRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceStoppedByUserRequests,
		prometheus.GaugeValue,
		instanceStoppedByUserRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceStoppedNoCapacityRequests,
		prometheus.GaugeValue,
		instanceStoppedNoCapacityRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedByPriceRequests,
		prometheus.GaugeValue,
		instanceTerminatedByPriceRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedByScheduleRequests,
		prometheus.GaugeValue,
		instanceTerminatedByScheduleRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedByServiceRequests,
		prometheus.GaugeValue,
		instanceTerminatedByServiceRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedByUserRequests,
		prometheus.GaugeValue,
		instanceTerminatedByUserRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedLaunchGroupConstraintRequests,
		prometheus.GaugeValue,
		instanceTerminatedLaunchGroupConstraintRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceTerminatedNoCapacityRequests,
		prometheus.GaugeValue,
		instanceTerminatedNoCapacityRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceLaunchGroupConstraintRequests,
		prometheus.GaugeValue,
		launchGroupConstraintRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceMarkedForStopRequests,
		prometheus.GaugeValue,
		markedForStopRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceMarkedForTerminationRequests,
		prometheus.GaugeValue,
		markedForTerminationRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceNotScheduledYetRequests,
		prometheus.GaugeValue,
		notScheduledYetRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstancePendingEvaluationRequests,
		prometheus.GaugeValue,
		pendingEvaluationRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstancePendingFulfillmentRequests,
		prometheus.GaugeValue,
		pendingFulfillmentRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstancePlacementGroupConstraintRequests,
		prometheus.GaugeValue,
		placementGroupConstraintRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstancePriceTooLowRequests,
		prometheus.GaugeValue,
		priceTooLowRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceRequestCanceledAndInstanceRunningRequests,
		prometheus.GaugeValue,
		requestCanceledAndInstanceRunningRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceScheduleExpiredRequests,
		prometheus.GaugeValue,
		scheduleExpiredRequests,
		region,
	)
	ch <- prometheus.MustNewConstMetric(
		c.spotInstanceSystemErrorRequests,
		prometheus.GaugeValue,
		systemErrorRequests,
		region,
	)
}

func init() {
	prometheus.MustRegister(NewCollector())
}

func main() {
	http.Handle(MetricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<title>AWS EC2 Spot Exporter</title>
			<body>
			<h1>AWS EC2 Spot Exporter</h1>
			<p><a href='` + MetricsPath + `'>Metrics</a></p>
			</html>`))
		if err != nil {
			log.Printf("Write failed: %v", err)
		}
	})

	log.Print("Starting EC2 Spot Exporter on port 9671...")
	log.Fatal(http.ListenAndServe(":9671", nil))
}
