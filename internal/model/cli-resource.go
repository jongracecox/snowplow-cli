/**
 * Copyright (c) 2013-present Snowplow Analytics Ltd.
 * All rights reserved.
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.0
 * located at https://docs.snowplow.io/limited-use-license-1.0
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */

package model

type CliResource[A any] struct {
	ApiVersion   string `yaml:"apiVersion" json:"apiVersion"`
	ResourceType string `yaml:"resourceType" json:"resourceType"`
	ResourceName string `yaml:"resourceName" json:"resourceName"`
	Data         A      `yaml:"data" json:"data"`
}
