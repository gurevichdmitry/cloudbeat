package compliance.cis_gcp.rules.cis_1_14

import data.cis_gcp.test_data
import data.compliance.policy.gcp.data_adapter
import data.lib.test
import future.keywords.if

type := "key-management"

subtype := "gcp-apikeys-key"

test_violation if {
	# fails when restrictions key doesn't exists
	eval_fail with input as test_data.generate_gcp_asset(
		type, subtype, {"data": {}},
		{},
	)

	# fails when restrictions key exists, but without at least 1 restriction
	eval_fail with input as test_data.generate_gcp_asset(
		type, subtype, {"data": {"restrictions": {"apiTargets": []}}},
		{},
	)

	# fails when google apis are restricted
	eval_fail with input as test_data.generate_gcp_asset(
		type, subtype, {"data": {"restrictions": {"apiTargets": [{"service": "cloudapis.googleapis.com"}]}}},
		{},
	)
}

test_pass if {
	eval_pass with input as test_data.generate_gcp_asset(
		type, subtype, {"data": {"restrictions": {"apiTargets": [{"service": "foo"}]}}},
		{},
	)
}

test_not_evaluated if {
	not_eval with input as test_data.not_eval_resource
}

eval_fail if {
	test.assert_fail(finding) with data.benchmark_data_adapter as data_adapter
}

eval_pass if {
	test.assert_pass(finding) with data.benchmark_data_adapter as data_adapter
}

not_eval if {
	not finding with data.benchmark_data_adapter as data_adapter
}
