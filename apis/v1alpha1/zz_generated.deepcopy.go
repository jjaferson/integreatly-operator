// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertingEmailAddresses) DeepCopyInto(out *AlertingEmailAddresses) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertingEmailAddresses.
func (in *AlertingEmailAddresses) DeepCopy() *AlertingEmailAddresses {
	if in == nil {
		return nil
	}
	out := new(AlertingEmailAddresses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Maintenance) DeepCopyInto(out *Maintenance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Maintenance.
func (in *Maintenance) DeepCopy() *Maintenance {
	if in == nil {
		return nil
	}
	out := new(Maintenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullSecretSpec) DeepCopyInto(out *PullSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullSecretSpec.
func (in *PullSecretSpec) DeepCopy() *PullSecretSpec {
	if in == nil {
		return nil
	}
	out := new(PullSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMI) DeepCopyInto(out *RHMI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMI.
func (in *RHMI) DeepCopy() *RHMI {
	if in == nil {
		return nil
	}
	out := new(RHMI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RHMI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfig) DeepCopyInto(out *RHMIConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfig.
func (in *RHMIConfig) DeepCopy() *RHMIConfig {
	if in == nil {
		return nil
	}
	out := new(RHMIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RHMIConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfigList) DeepCopyInto(out *RHMIConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RHMIConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfigList.
func (in *RHMIConfigList) DeepCopy() *RHMIConfigList {
	if in == nil {
		return nil
	}
	out := new(RHMIConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RHMIConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfigSpec) DeepCopyInto(out *RHMIConfigSpec) {
	*out = *in
	in.Upgrade.DeepCopyInto(&out.Upgrade)
	out.Maintenance = in.Maintenance
	out.Backup = in.Backup
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfigSpec.
func (in *RHMIConfigSpec) DeepCopy() *RHMIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(RHMIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfigStatus) DeepCopyInto(out *RHMIConfigStatus) {
	*out = *in
	out.Maintenance = in.Maintenance
	in.Upgrade.DeepCopyInto(&out.Upgrade)
	if in.UpgradeAvailable != nil {
		in, out := &in.UpgradeAvailable, &out.UpgradeAvailable
		*out = new(UpgradeAvailable)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfigStatus.
func (in *RHMIConfigStatus) DeepCopy() *RHMIConfigStatus {
	if in == nil {
		return nil
	}
	out := new(RHMIConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfigStatusMaintenance) DeepCopyInto(out *RHMIConfigStatusMaintenance) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfigStatusMaintenance.
func (in *RHMIConfigStatusMaintenance) DeepCopy() *RHMIConfigStatusMaintenance {
	if in == nil {
		return nil
	}
	out := new(RHMIConfigStatusMaintenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIConfigStatusUpgrade) DeepCopyInto(out *RHMIConfigStatusUpgrade) {
	*out = *in
	if in.Scheduled != nil {
		in, out := &in.Scheduled, &out.Scheduled
		*out = new(UpgradeSchedule)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIConfigStatusUpgrade.
func (in *RHMIConfigStatusUpgrade) DeepCopy() *RHMIConfigStatusUpgrade {
	if in == nil {
		return nil
	}
	out := new(RHMIConfigStatusUpgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIList) DeepCopyInto(out *RHMIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RHMI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIList.
func (in *RHMIList) DeepCopy() *RHMIList {
	if in == nil {
		return nil
	}
	out := new(RHMIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RHMIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIProductStatus) DeepCopyInto(out *RHMIProductStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIProductStatus.
func (in *RHMIProductStatus) DeepCopy() *RHMIProductStatus {
	if in == nil {
		return nil
	}
	out := new(RHMIProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMISpec) DeepCopyInto(out *RHMISpec) {
	*out = *in
	out.PullSecret = in.PullSecret
	out.AlertingEmailAddresses = in.AlertingEmailAddresses
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMISpec.
func (in *RHMISpec) DeepCopy() *RHMISpec {
	if in == nil {
		return nil
	}
	out := new(RHMISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIStageStatus) DeepCopyInto(out *RHMIStageStatus) {
	*out = *in
	if in.Products != nil {
		in, out := &in.Products, &out.Products
		*out = make(map[ProductName]RHMIProductStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIStageStatus.
func (in *RHMIStageStatus) DeepCopy() *RHMIStageStatus {
	if in == nil {
		return nil
	}
	out := new(RHMIStageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHMIStatus) DeepCopyInto(out *RHMIStatus) {
	*out = *in
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make(map[StageName]RHMIStageStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHMIStatus.
func (in *RHMIStatus) DeepCopy() *RHMIStatus {
	if in == nil {
		return nil
	}
	out := new(RHMIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upgrade) DeepCopyInto(out *Upgrade) {
	*out = *in
	if in.WaitForMaintenance != nil {
		in, out := &in.WaitForMaintenance, &out.WaitForMaintenance
		*out = new(bool)
		**out = **in
	}
	if in.NotBeforeDays != nil {
		in, out := &in.NotBeforeDays, &out.NotBeforeDays
		*out = new(int)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upgrade.
func (in *Upgrade) DeepCopy() *Upgrade {
	if in == nil {
		return nil
	}
	out := new(Upgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeAvailable) DeepCopyInto(out *UpgradeAvailable) {
	*out = *in
	in.AvailableAt.DeepCopyInto(&out.AvailableAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeAvailable.
func (in *UpgradeAvailable) DeepCopy() *UpgradeAvailable {
	if in == nil {
		return nil
	}
	out := new(UpgradeAvailable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeSchedule) DeepCopyInto(out *UpgradeSchedule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeSchedule.
func (in *UpgradeSchedule) DeepCopy() *UpgradeSchedule {
	if in == nil {
		return nil
	}
	out := new(UpgradeSchedule)
	in.DeepCopyInto(out)
	return out
}
