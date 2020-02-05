# EC2 Spot Instance Exporter
A prometheus exporter providing metrics for AWS EC2 Spot Instance usage. The exporter aims to provide visibility into Spot Instance requests which can be used to improve uptime.

## Building
```
make
```
## Running
```
./ec2-spot-exporter
```
## Metrics
| Name | Description |
| ---- | ---- |
| ec2_spot_instance_requests | Spot instance requests count |
| ec2_spot_instance_az_group_contraint_requests | Spot instance requests with az-group-constraint status count |
| ec2_spot_instance_bad_parameters_requests | Spot instance requests with az-group-constraint status count |
| ec2_spot_instance_canceled_before_fulfillment_requests | Spot instance requests with canceled-before-fulfillment status count |
| ec2_spot_instance_capacity_not_available_requests | Spot instance requests with capacity-not-available status count |
| ec2_spot_instance_constraint_not_fulfillable_requests | Spot instance requests with constraint-not-fulfillable status count |
| ec2_spot_instance_fulfilled_requests | Spot instance requests with fulfilled status count |
| ec2_spot_instance_stopped_by_price_requests | Spot instance requests with instance-stopped-by-price status count |
| ec2_spot_instance_stopped_by_user_requests | Spot instance requests with instance-stopped-by-user status count |
| ec2_spot_instance_stopped_no_capacity_requests | Spot instance requests with instance-stopped-no-capacity status count |
| ec2_spot_instance_terminated_by_price_requests | Spot instance requests with instance-terminated-by-price status count |
| ec2_spot_instance_terminated_by_schedule_requests | Spot instance requests with instance-terminated-by-schedule status count |
| ec2_spot_instance_terminated_by_service_requests | Spot instance requests with instance-terminated-by-service status count |
| ec2_spot_instance_terminated_by_user_requests | Spot instance requests with instance-terminated-by-user status count |
| ec2_spot_instance_terminated_launch_group_constraint_requests | Spot instance requests with instance-terminated-launch-group-constraint status count |
| ec2_spot_instance_terminated_no_capacity_requests | Spot instance requests with instance-terminated-no-capacity status count |
| ec2_spot_instance_launch_group_constraint_requests | Spot instance requests with launch-group-constraint status count |
| ec2_spot_instance_marked_for_stop_requests | Spot instance requests with marked-for-stop status count |
| ec2_spot_instance_marked_for_termination_requests | Spot instance requests with marked-for-termination status count |
| ec2_spot_instance_not_scheduled_yet_requests | Spot instance requests with not-scheduled-yet status count |
| ec2_spot_instance_pending_evaluation_requests | Spot instance requests with pending-evaluation status count |
| ec2_spot_instance_pending_fulfillment_requests | Spot instance requests with pending-fulfillment status count |
| ec2_spot_instance_placement_group_constraint_requests | Spot instance requests with placement-group-constraint status count |
| ec2_spot_instance_price_too_low_requests | Spot instance requests with price-too-low status count |
| ec2_spot_instance_request_canceled_and_instance_running_requests | Spot instance requests with request-canceled-and-instance-running status count |
| ec2_spot_instance_schedule_expired_requests | Spot instance requests with schedule_expired status count |
| ec2_spot_instance_system_error_requests | Spot instance requests with system-error status count |