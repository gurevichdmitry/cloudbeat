package compliance.cis_k8s.rules.cis_5_2_6

import data.compliance.policy.kube_api.minimize_admission as audit

finding := audit.finding("allowPrivilegeEscalation")
