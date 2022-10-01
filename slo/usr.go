package slo

type Data struct {
	// Name
	// Description
	// uuid
	// Metric : response_code, duration, response_flags, grpc_status
	// Calculation : % of traffic, percentiles, count in a time frame,
	//               derivative vs actual calculation
	// Watch and Warning configuration, analogous to the National Weather Service definition;
	// Filter : traffic,region,zone,sub-zone,route_name
	// Traffic : ingress,egress,ping
	// Enabled : true,false
	// Window  : inclusive or exclusive time frame for applicability of SLO
	// Notification : true false
	// Silence notifications
	// Notification Filter : watches and warnings
}
