package compliance.cis_azure.rules.cis_3_2

import data.cis_azure.test_data
import data.compliance.policy.azure.data_adapter
import data.lib.test
import future.keywords.if

test_violation if {
	# fail if managed by user
	eval_fail with input as test_data.generate_storage_account_with_property("encryption", {"requireInfrastructureEncryption": false})
	eval_fail with input as test_data.generate_storage_account_with_property("encryption", {})
}

test_pass if {
	# pass if not managed by user
	eval_pass with input as test_data.generate_storage_account_with_property("encryption", {"requireInfrastructureEncryption": true})
}

test_not_evaluated if {
	not_eval with input as test_data.not_eval_non_exist_type
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
