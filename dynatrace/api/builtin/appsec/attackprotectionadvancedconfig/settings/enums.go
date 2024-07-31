/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package attackprotectionadvancedconfig

type AttackType string

var AttackTypes = struct {
	Any           AttackType
	CmdInjection  AttackType
	JndiInjection AttackType
	SqlInjection  AttackType
	Ssrf          AttackType
}{
	"ANY",
	"CMD_INJECTION",
	"JNDI_INJECTION",
	"SQL_INJECTION",
	"SSRF",
}

type BlockingStrategy string

var BlockingStrategies = struct {
	Block   BlockingStrategy
	Monitor BlockingStrategy
	Off     BlockingStrategy
}{
	"BLOCK",
	"MONITOR",
	"OFF",
}

type ResourceAttributeValueMatcher string

var ResourceAttributeValueMatchers = struct {
	Contains         ResourceAttributeValueMatcher
	DoesNotContain   ResourceAttributeValueMatcher
	DoesNotEndWith   ResourceAttributeValueMatcher
	DoesNotExist     ResourceAttributeValueMatcher
	DoesNotStartWith ResourceAttributeValueMatcher
	EndsWith         ResourceAttributeValueMatcher
	Equals           ResourceAttributeValueMatcher
	Exists           ResourceAttributeValueMatcher
	NotEquals        ResourceAttributeValueMatcher
	StartsWith       ResourceAttributeValueMatcher
}{
	"CONTAINS",
	"DOES_NOT_CONTAIN",
	"DOES_NOT_END_WITH",
	"DOES_NOT_EXIST",
	"DOES_NOT_START_WITH",
	"ENDS_WITH",
	"EQUALS",
	"EXISTS",
	"NOT_EQUALS",
	"STARTS_WITH",
}
