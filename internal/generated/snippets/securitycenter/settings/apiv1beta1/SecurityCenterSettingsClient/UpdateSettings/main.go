// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START securitycenter_generated_securitycenter_settings_apiv1beta1_SecurityCenterSettingsClient_UpdateSettings]

package main

import (
	"context"

	settings "cloud.google.com/go/securitycenter/settings/apiv1beta1"
	settingspb "google.golang.org/genproto/googleapis/cloud/securitycenter/settings/v1beta1"
)

func main() {
	// import settingspb "google.golang.org/genproto/googleapis/cloud/securitycenter/settings/v1beta1"

	ctx := context.Background()
	c, err := settings.NewSecurityCenterSettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &settingspb.UpdateSettingsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END securitycenter_generated_securitycenter_settings_apiv1beta1_SecurityCenterSettingsClient_UpdateSettings]
