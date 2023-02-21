package ghid

const (
	// TypeAddedToProjectEvent is constant for a type of AddedToProjectEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#addedtoprojectevent.
	TypeAddedToProjectEvent = "AddedToProjectEvent"

	// TypeApp is constant for a type of App node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#app.
	TypeApp = "App"

	// TypeAssignedEvent is constant for a type of AssignedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#assignedevent.
	TypeAssignedEvent = "AssignedEvent"

	// TypeAutoMergeDisabledEvent is constant for a type of AutoMergeDisabledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#automergedisabledevent.
	TypeAutoMergeDisabledEvent = "AutoMergeDisabledEvent"

	// TypeAutoMergeEnabledEvent is constant for a type of AutoMergeEnabledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#automergeenabledevent.
	TypeAutoMergeEnabledEvent = "AutoMergeEnabledEvent"

	// TypeAutoRebaseEnabledEvent is constant for a type of AutoRebaseEnabledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#autorebaseenabledevent.
	TypeAutoRebaseEnabledEvent = "AutoRebaseEnabledEvent"

	// TypeAutoSquashEnabledEvent is constant for a type of AutoSquashEnabledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#autosquashenabledevent.
	TypeAutoSquashEnabledEvent = "AutoSquashEnabledEvent"

	// TypeAutomaticBaseChangeFailedEvent is constant for a type of AutomaticBaseChangeFailedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#automaticbasechangefailedevent.
	TypeAutomaticBaseChangeFailedEvent = "AutomaticBaseChangeFailedEvent"

	// TypeAutomaticBaseChangeSucceededEvent is constant for a type of AutomaticBaseChangeSucceededEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#automaticbasechangesucceededevent.
	TypeAutomaticBaseChangeSucceededEvent = "AutomaticBaseChangeSucceededEvent"

	// TypeBaseRefChangedEvent is constant for a type of BaseRefChangedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#baserefchangedevent.
	TypeBaseRefChangedEvent = "BaseRefChangedEvent"

	// TypeBaseRefDeletedEvent is constant for a type of BaseRefDeletedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#baserefdeletedevent.
	TypeBaseRefDeletedEvent = "BaseRefDeletedEvent"

	// TypeBaseRefForcePushedEvent is constant for a type of BaseRefForcePushedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#baserefforcepushedevent.
	TypeBaseRefForcePushedEvent = "BaseRefForcePushedEvent"

	// TypeBlob is constant for a type of Blob node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#blob.
	TypeBlob = "Blob"

	// TypeBot is constant for a type of Bot node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#bot.
	TypeBot = "Bot"

	// TypeBranchProtectionRule is constant for a type of BranchProtectionRule node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#branchprotectionrule.
	TypeBranchProtectionRule = "BranchProtectionRule"

	// TypeBypassForcePushAllowance is constant for a type of BypassForcePushAllowance node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#bypassforcepushallowance.
	TypeBypassForcePushAllowance = "BypassForcePushAllowance"

	// TypeBypassPullRequestAllowance is constant for a type of BypassPullRequestAllowance node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#bypasspullrequestallowance.
	TypeBypassPullRequestAllowance = "BypassPullRequestAllowance"

	// TypeCWE is constant for a type of CWE node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#cwe.
	TypeCWE = "CWE"

	// TypeCheckRun is constant for a type of CheckRun node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#checkrun.
	TypeCheckRun = "CheckRun"

	// TypeCheckSuite is constant for a type of CheckSuite node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#checksuite.
	TypeCheckSuite = "CheckSuite"

	// TypeClosedEvent is constant for a type of ClosedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#closedevent.
	TypeClosedEvent = "ClosedEvent"

	// TypeCodeOfConduct is constant for a type of CodeOfConduct node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#codeofconduct.
	TypeCodeOfConduct = "CodeOfConduct"

	// TypeCommentDeletedEvent is constant for a type of CommentDeletedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#commentdeletedevent.
	TypeCommentDeletedEvent = "CommentDeletedEvent"

	// TypeCommit is constant for a type of Commit node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#commit.
	TypeCommit = "Commit"

	// TypeCommitComment is constant for a type of CommitComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#commitcomment.
	TypeCommitComment = "CommitComment"

	// TypeCommitCommentThread is constant for a type of CommitCommentThread node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#commitcommentthread.
	TypeCommitCommentThread = "CommitCommentThread"

	// TypeComparison is constant for a type of Comparison node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#comparison.
	TypeComparison = "Comparison"

	// TypeConnectedEvent is constant for a type of ConnectedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#connectedevent.
	TypeConnectedEvent = "ConnectedEvent"

	// TypeConvertToDraftEvent is constant for a type of ConvertToDraftEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#converttodraftevent.
	TypeConvertToDraftEvent = "ConvertToDraftEvent"

	// TypeConvertedNoteToIssueEvent is constant for a type of ConvertedNoteToIssueEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#convertednotetoissueevent.
	TypeConvertedNoteToIssueEvent = "ConvertedNoteToIssueEvent"

	// TypeConvertedToDiscussionEvent is constant for a type of ConvertedToDiscussionEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#convertedtodiscussionevent.
	TypeConvertedToDiscussionEvent = "ConvertedToDiscussionEvent"

	// TypeCrossReferencedEvent is constant for a type of CrossReferencedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#crossreferencedevent.
	TypeCrossReferencedEvent = "CrossReferencedEvent"

	// TypeDemilestonedEvent is constant for a type of DemilestonedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#demilestonedevent.
	TypeDemilestonedEvent = "DemilestonedEvent"

	// TypeDeployKey is constant for a type of DeployKey node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deploykey.
	TypeDeployKey = "DeployKey"

	// TypeDeployedEvent is constant for a type of DeployedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deployedevent.
	TypeDeployedEvent = "DeployedEvent"

	// TypeDeployment is constant for a type of Deployment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deployment.
	TypeDeployment = "Deployment"

	// TypeDeploymentEnvironmentChangedEvent is constant for a type of DeploymentEnvironmentChangedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deploymentenvironmentchangedevent.
	TypeDeploymentEnvironmentChangedEvent = "DeploymentEnvironmentChangedEvent"

	// TypeDeploymentReview is constant for a type of DeploymentReview node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deploymentreview.
	TypeDeploymentReview = "DeploymentReview"

	// TypeDeploymentStatus is constant for a type of DeploymentStatus node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#deploymentstatus.
	TypeDeploymentStatus = "DeploymentStatus"

	// TypeDisconnectedEvent is constant for a type of DisconnectedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#disconnectedevent.
	TypeDisconnectedEvent = "DisconnectedEvent"

	// TypeDiscussion is constant for a type of Discussion node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#discussion.
	TypeDiscussion = "Discussion"

	// TypeDiscussionCategory is constant for a type of DiscussionCategory node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#discussioncategory.
	TypeDiscussionCategory = "DiscussionCategory"

	// TypeDiscussionComment is constant for a type of DiscussionComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#discussioncomment.
	TypeDiscussionComment = "DiscussionComment"

	// TypeDiscussionPoll is constant for a type of DiscussionPoll node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#discussionpoll.
	TypeDiscussionPoll = "DiscussionPoll"

	// TypeDiscussionPollOption is constant for a type of DiscussionPollOption node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#discussionpolloption.
	TypeDiscussionPollOption = "DiscussionPollOption"

	// TypeDraftIssue is constant for a type of DraftIssue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#draftissue.
	TypeDraftIssue = "DraftIssue"

	// TypeEnterprise is constant for a type of Enterprise node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterprise.
	TypeEnterprise = "Enterprise"

	// TypeEnterpriseAdministratorInvitation is constant for a type of EnterpriseAdministratorInvitation node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseadministratorinvitation.
	TypeEnterpriseAdministratorInvitation = "EnterpriseAdministratorInvitation"

	// TypeEnterpriseIdentityProvider is constant for a type of EnterpriseIdentityProvider node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseidentityprovider.
	TypeEnterpriseIdentityProvider = "EnterpriseIdentityProvider"

	// TypeEnterpriseRepositoryInfo is constant for a type of EnterpriseRepositoryInfo node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriserepositoryinfo.
	TypeEnterpriseRepositoryInfo = "EnterpriseRepositoryInfo"

	// TypeEnterpriseServerInstallation is constant for a type of EnterpriseServerInstallation node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseserverinstallation.
	TypeEnterpriseServerInstallation = "EnterpriseServerInstallation"

	// TypeEnterpriseServerUserAccount is constant for a type of EnterpriseServerUserAccount node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseserveruseraccount.
	TypeEnterpriseServerUserAccount = "EnterpriseServerUserAccount"

	// TypeEnterpriseServerUserAccountEmail is constant for a type of EnterpriseServerUserAccountEmail node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseserveruseraccountemail.
	TypeEnterpriseServerUserAccountEmail = "EnterpriseServerUserAccountEmail"

	// TypeEnterpriseServerUserAccountsUpload is constant for a type of EnterpriseServerUserAccountsUpload node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseserveruseraccountsupload.
	TypeEnterpriseServerUserAccountsUpload = "EnterpriseServerUserAccountsUpload"

	// TypeEnterpriseUserAccount is constant for a type of EnterpriseUserAccount node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#enterpriseuseraccount.
	TypeEnterpriseUserAccount = "EnterpriseUserAccount"

	// TypeEnvironment is constant for a type of Environment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#environment.
	TypeEnvironment = "Environment"

	// TypeExternalIdentity is constant for a type of ExternalIdentity node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#externalidentity.
	TypeExternalIdentity = "ExternalIdentity"

	// TypeGist is constant for a type of Gist node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#gist.
	TypeGist = "Gist"

	// TypeGistComment is constant for a type of GistComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#gistcomment.
	TypeGistComment = "GistComment"

	// TypeHeadRefDeletedEvent is constant for a type of HeadRefDeletedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#headrefdeletedevent.
	TypeHeadRefDeletedEvent = "HeadRefDeletedEvent"

	// TypeHeadRefForcePushedEvent is constant for a type of HeadRefForcePushedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#headrefforcepushedevent.
	TypeHeadRefForcePushedEvent = "HeadRefForcePushedEvent"

	// TypeHeadRefRestoredEvent is constant for a type of HeadRefRestoredEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#headrefrestoredevent.
	TypeHeadRefRestoredEvent = "HeadRefRestoredEvent"

	// TypeIpAllowListEntry is constant for a type of IpAllowListEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#ipallowlistentry.
	TypeIpAllowListEntry = "IpAllowListEntry"

	// TypeIssue is constant for a type of Issue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#issue.
	TypeIssue = "Issue"

	// TypeIssueComment is constant for a type of IssueComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#issuecomment.
	TypeIssueComment = "IssueComment"

	// TypeLabel is constant for a type of Label node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#label.
	TypeLabel = "Label"

	// TypeLabeledEvent is constant for a type of LabeledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#labeledevent.
	TypeLabeledEvent = "LabeledEvent"

	// TypeLanguage is constant for a type of Language node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#language.
	TypeLanguage = "Language"

	// TypeLicense is constant for a type of License node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#license.
	TypeLicense = "License"

	// TypeLinkedBranch is constant for a type of LinkedBranch node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#linkedbranch.
	TypeLinkedBranch = "LinkedBranch"

	// TypeLockedEvent is constant for a type of LockedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#lockedevent.
	TypeLockedEvent = "LockedEvent"

	// TypeMannequin is constant for a type of Mannequin node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#mannequin.
	TypeMannequin = "Mannequin"

	// TypeMarkedAsDuplicateEvent is constant for a type of MarkedAsDuplicateEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#markedasduplicateevent.
	TypeMarkedAsDuplicateEvent = "MarkedAsDuplicateEvent"

	// TypeMarketplaceCategory is constant for a type of MarketplaceCategory node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#marketplacecategory.
	TypeMarketplaceCategory = "MarketplaceCategory"

	// TypeMarketplaceListing is constant for a type of MarketplaceListing node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#marketplacelisting.
	TypeMarketplaceListing = "MarketplaceListing"

	// TypeMembersCanDeleteReposClearAuditEntry is constant for a type of MembersCanDeleteReposClearAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#memberscandeletereposclearauditentry.
	TypeMembersCanDeleteReposClearAuditEntry = "MembersCanDeleteReposClearAuditEntry"

	// TypeMembersCanDeleteReposDisableAuditEntry is constant for a type of MembersCanDeleteReposDisableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#memberscandeletereposdisableauditentry.
	TypeMembersCanDeleteReposDisableAuditEntry = "MembersCanDeleteReposDisableAuditEntry"

	// TypeMembersCanDeleteReposEnableAuditEntry is constant for a type of MembersCanDeleteReposEnableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#memberscandeletereposenableauditentry.
	TypeMembersCanDeleteReposEnableAuditEntry = "MembersCanDeleteReposEnableAuditEntry"

	// TypeMentionedEvent is constant for a type of MentionedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#mentionedevent.
	TypeMentionedEvent = "MentionedEvent"

	// TypeMergedEvent is constant for a type of MergedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#mergedevent.
	TypeMergedEvent = "MergedEvent"

	// TypeMigrationSource is constant for a type of MigrationSource node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#migrationsource.
	TypeMigrationSource = "MigrationSource"

	// TypeMilestone is constant for a type of Milestone node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#milestone.
	TypeMilestone = "Milestone"

	// TypeMilestonedEvent is constant for a type of MilestonedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#milestonedevent.
	TypeMilestonedEvent = "MilestonedEvent"

	// TypeMovedColumnsInProjectEvent is constant for a type of MovedColumnsInProjectEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#movedcolumnsinprojectevent.
	TypeMovedColumnsInProjectEvent = "MovedColumnsInProjectEvent"

	// TypeOIDCProvider is constant for a type of OIDCProvider node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#oidcprovider.
	TypeOIDCProvider = "OIDCProvider"

	// TypeOauthApplicationCreateAuditEntry is constant for a type of OauthApplicationCreateAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#oauthapplicationcreateauditentry.
	TypeOauthApplicationCreateAuditEntry = "OauthApplicationCreateAuditEntry"

	// TypeOrgAddBillingManagerAuditEntry is constant for a type of OrgAddBillingManagerAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgaddbillingmanagerauditentry.
	TypeOrgAddBillingManagerAuditEntry = "OrgAddBillingManagerAuditEntry"

	// TypeOrgAddMemberAuditEntry is constant for a type of OrgAddMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgaddmemberauditentry.
	TypeOrgAddMemberAuditEntry = "OrgAddMemberAuditEntry"

	// TypeOrgBlockUserAuditEntry is constant for a type of OrgBlockUserAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgblockuserauditentry.
	TypeOrgBlockUserAuditEntry = "OrgBlockUserAuditEntry"

	// TypeOrgConfigDisableCollaboratorsOnlyAuditEntry is constant for a type of OrgConfigDisableCollaboratorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgconfigdisablecollaboratorsonlyauditentry.
	TypeOrgConfigDisableCollaboratorsOnlyAuditEntry = "OrgConfigDisableCollaboratorsOnlyAuditEntry"

	// TypeOrgConfigEnableCollaboratorsOnlyAuditEntry is constant for a type of OrgConfigEnableCollaboratorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgconfigenablecollaboratorsonlyauditentry.
	TypeOrgConfigEnableCollaboratorsOnlyAuditEntry = "OrgConfigEnableCollaboratorsOnlyAuditEntry"

	// TypeOrgCreateAuditEntry is constant for a type of OrgCreateAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgcreateauditentry.
	TypeOrgCreateAuditEntry = "OrgCreateAuditEntry"

	// TypeOrgDisableOauthAppRestrictionsAuditEntry is constant for a type of OrgDisableOauthAppRestrictionsAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgdisableoauthapprestrictionsauditentry.
	TypeOrgDisableOauthAppRestrictionsAuditEntry = "OrgDisableOauthAppRestrictionsAuditEntry"

	// TypeOrgDisableSamlAuditEntry is constant for a type of OrgDisableSamlAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgdisablesamlauditentry.
	TypeOrgDisableSamlAuditEntry = "OrgDisableSamlAuditEntry"

	// TypeOrgDisableTwoFactorRequirementAuditEntry is constant for a type of OrgDisableTwoFactorRequirementAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgdisabletwofactorrequirementauditentry.
	TypeOrgDisableTwoFactorRequirementAuditEntry = "OrgDisableTwoFactorRequirementAuditEntry"

	// TypeOrgEnableOauthAppRestrictionsAuditEntry is constant for a type of OrgEnableOauthAppRestrictionsAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgenableoauthapprestrictionsauditentry.
	TypeOrgEnableOauthAppRestrictionsAuditEntry = "OrgEnableOauthAppRestrictionsAuditEntry"

	// TypeOrgEnableSamlAuditEntry is constant for a type of OrgEnableSamlAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgenablesamlauditentry.
	TypeOrgEnableSamlAuditEntry = "OrgEnableSamlAuditEntry"

	// TypeOrgEnableTwoFactorRequirementAuditEntry is constant for a type of OrgEnableTwoFactorRequirementAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgenabletwofactorrequirementauditentry.
	TypeOrgEnableTwoFactorRequirementAuditEntry = "OrgEnableTwoFactorRequirementAuditEntry"

	// TypeOrgInviteMemberAuditEntry is constant for a type of OrgInviteMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orginvitememberauditentry.
	TypeOrgInviteMemberAuditEntry = "OrgInviteMemberAuditEntry"

	// TypeOrgInviteToBusinessAuditEntry is constant for a type of OrgInviteToBusinessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orginvitetobusinessauditentry.
	TypeOrgInviteToBusinessAuditEntry = "OrgInviteToBusinessAuditEntry"

	// TypeOrgOauthAppAccessApprovedAuditEntry is constant for a type of OrgOauthAppAccessApprovedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgoauthappaccessapprovedauditentry.
	TypeOrgOauthAppAccessApprovedAuditEntry = "OrgOauthAppAccessApprovedAuditEntry"

	// TypeOrgOauthAppAccessDeniedAuditEntry is constant for a type of OrgOauthAppAccessDeniedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgoauthappaccessdeniedauditentry.
	TypeOrgOauthAppAccessDeniedAuditEntry = "OrgOauthAppAccessDeniedAuditEntry"

	// TypeOrgOauthAppAccessRequestedAuditEntry is constant for a type of OrgOauthAppAccessRequestedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgoauthappaccessrequestedauditentry.
	TypeOrgOauthAppAccessRequestedAuditEntry = "OrgOauthAppAccessRequestedAuditEntry"

	// TypeOrgRemoveBillingManagerAuditEntry is constant for a type of OrgRemoveBillingManagerAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgremovebillingmanagerauditentry.
	TypeOrgRemoveBillingManagerAuditEntry = "OrgRemoveBillingManagerAuditEntry"

	// TypeOrgRemoveMemberAuditEntry is constant for a type of OrgRemoveMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgremovememberauditentry.
	TypeOrgRemoveMemberAuditEntry = "OrgRemoveMemberAuditEntry"

	// TypeOrgRemoveOutsideCollaboratorAuditEntry is constant for a type of OrgRemoveOutsideCollaboratorAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgremoveoutsidecollaboratorauditentry.
	TypeOrgRemoveOutsideCollaboratorAuditEntry = "OrgRemoveOutsideCollaboratorAuditEntry"

	// TypeOrgRestoreMemberAuditEntry is constant for a type of OrgRestoreMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgrestorememberauditentry.
	TypeOrgRestoreMemberAuditEntry = "OrgRestoreMemberAuditEntry"

	// TypeOrgUnblockUserAuditEntry is constant for a type of OrgUnblockUserAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgunblockuserauditentry.
	TypeOrgUnblockUserAuditEntry = "OrgUnblockUserAuditEntry"

	// TypeOrgUpdateDefaultRepositoryPermissionAuditEntry is constant for a type of OrgUpdateDefaultRepositoryPermissionAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgupdatedefaultrepositorypermissionauditentry.
	TypeOrgUpdateDefaultRepositoryPermissionAuditEntry = "OrgUpdateDefaultRepositoryPermissionAuditEntry"

	// TypeOrgUpdateMemberAuditEntry is constant for a type of OrgUpdateMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgupdatememberauditentry.
	TypeOrgUpdateMemberAuditEntry = "OrgUpdateMemberAuditEntry"

	// TypeOrgUpdateMemberRepositoryCreationPermissionAuditEntry is constant for a type of OrgUpdateMemberRepositoryCreationPermissionAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgupdatememberrepositorycreationpermissionauditentry.
	TypeOrgUpdateMemberRepositoryCreationPermissionAuditEntry = "OrgUpdateMemberRepositoryCreationPermissionAuditEntry"

	// TypeOrgUpdateMemberRepositoryInvitationPermissionAuditEntry is constant for a type of OrgUpdateMemberRepositoryInvitationPermissionAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#orgupdatememberrepositoryinvitationpermissionauditentry.
	TypeOrgUpdateMemberRepositoryInvitationPermissionAuditEntry = "OrgUpdateMemberRepositoryInvitationPermissionAuditEntry"

	// TypeOrganization is constant for a type of Organization node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#organization.
	TypeOrganization = "Organization"

	// TypeOrganizationIdentityProvider is constant for a type of OrganizationIdentityProvider node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#organizationidentityprovider.
	TypeOrganizationIdentityProvider = "OrganizationIdentityProvider"

	// TypeOrganizationInvitation is constant for a type of OrganizationInvitation node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#organizationinvitation.
	TypeOrganizationInvitation = "OrganizationInvitation"

	// TypeOrganizationMigration is constant for a type of OrganizationMigration node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#organizationmigration.
	TypeOrganizationMigration = "OrganizationMigration"

	// TypePackage is constant for a type of Package node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#package.
	TypePackage = "Package"

	// TypePackageFile is constant for a type of PackageFile node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#packagefile.
	TypePackageFile = "PackageFile"

	// TypePackageTag is constant for a type of PackageTag node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#packagetag.
	TypePackageTag = "PackageTag"

	// TypePackageVersion is constant for a type of PackageVersion node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#packageversion.
	TypePackageVersion = "PackageVersion"

	// TypePinnedDiscussion is constant for a type of PinnedDiscussion node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pinneddiscussion.
	TypePinnedDiscussion = "PinnedDiscussion"

	// TypePinnedEvent is constant for a type of PinnedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pinnedevent.
	TypePinnedEvent = "PinnedEvent"

	// TypePinnedIssue is constant for a type of PinnedIssue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pinnedissue.
	TypePinnedIssue = "PinnedIssue"

	// TypePrivateRepositoryForkingDisableAuditEntry is constant for a type of PrivateRepositoryForkingDisableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#privaterepositoryforkingdisableauditentry.
	TypePrivateRepositoryForkingDisableAuditEntry = "PrivateRepositoryForkingDisableAuditEntry"

	// TypePrivateRepositoryForkingEnableAuditEntry is constant for a type of PrivateRepositoryForkingEnableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#privaterepositoryforkingenableauditentry.
	TypePrivateRepositoryForkingEnableAuditEntry = "PrivateRepositoryForkingEnableAuditEntry"

	// TypeProject is constant for a type of Project node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#project.
	TypeProject = "Project"

	// TypeProjectCard is constant for a type of ProjectCard node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectcard.
	TypeProjectCard = "ProjectCard"

	// TypeProjectColumn is constant for a type of ProjectColumn node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectcolumn.
	TypeProjectColumn = "ProjectColumn"

	// TypeProjectV2 is constant for a type of ProjectV2 node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2.
	TypeProjectV2 = "ProjectV2"

	// TypeProjectV2Field is constant for a type of ProjectV2Field node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2field.
	TypeProjectV2Field = "ProjectV2Field"

	// TypeProjectV2Item is constant for a type of ProjectV2Item node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2item.
	TypeProjectV2Item = "ProjectV2Item"

	// TypeProjectV2ItemFieldDateValue is constant for a type of ProjectV2ItemFieldDateValue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2itemfielddatevalue.
	TypeProjectV2ItemFieldDateValue = "ProjectV2ItemFieldDateValue"

	// TypeProjectV2ItemFieldIterationValue is constant for a type of ProjectV2ItemFieldIterationValue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2itemfielditerationvalue.
	TypeProjectV2ItemFieldIterationValue = "ProjectV2ItemFieldIterationValue"

	// TypeProjectV2ItemFieldNumberValue is constant for a type of ProjectV2ItemFieldNumberValue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2itemfieldnumbervalue.
	TypeProjectV2ItemFieldNumberValue = "ProjectV2ItemFieldNumberValue"

	// TypeProjectV2ItemFieldSingleSelectValue is constant for a type of ProjectV2ItemFieldSingleSelectValue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2itemfieldsingleselectvalue.
	TypeProjectV2ItemFieldSingleSelectValue = "ProjectV2ItemFieldSingleSelectValue"

	// TypeProjectV2ItemFieldTextValue is constant for a type of ProjectV2ItemFieldTextValue node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2itemfieldtextvalue.
	TypeProjectV2ItemFieldTextValue = "ProjectV2ItemFieldTextValue"

	// TypeProjectV2IterationField is constant for a type of ProjectV2IterationField node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2iterationfield.
	TypeProjectV2IterationField = "ProjectV2IterationField"

	// TypeProjectV2SingleSelectField is constant for a type of ProjectV2SingleSelectField node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2singleselectfield.
	TypeProjectV2SingleSelectField = "ProjectV2SingleSelectField"

	// TypeProjectV2View is constant for a type of ProjectV2View node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2view.
	TypeProjectV2View = "ProjectV2View"

	// TypeProjectV2Workflow is constant for a type of ProjectV2Workflow node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#projectv2workflow.
	TypeProjectV2Workflow = "ProjectV2Workflow"

	// TypePublicKey is constant for a type of PublicKey node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#publickey.
	TypePublicKey = "PublicKey"

	// TypePullRequest is constant for a type of PullRequest node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequest.
	TypePullRequest = "PullRequest"

	// TypePullRequestCommit is constant for a type of PullRequestCommit node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestcommit.
	TypePullRequestCommit = "PullRequestCommit"

	// TypePullRequestCommitCommentThread is constant for a type of PullRequestCommitCommentThread node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestcommitcommentthread.
	TypePullRequestCommitCommentThread = "PullRequestCommitCommentThread"

	// TypePullRequestReview is constant for a type of PullRequestReview node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestreview.
	TypePullRequestReview = "PullRequestReview"

	// TypePullRequestReviewComment is constant for a type of PullRequestReviewComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewcomment.
	TypePullRequestReviewComment = "PullRequestReviewComment"

	// TypePullRequestReviewThread is constant for a type of PullRequestReviewThread node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewthread.
	TypePullRequestReviewThread = "PullRequestReviewThread"

	// TypePullRequestThread is constant for a type of PullRequestThread node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pullrequestthread.
	TypePullRequestThread = "PullRequestThread"

	// TypePush is constant for a type of Push node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#push.
	TypePush = "Push"

	// TypePushAllowance is constant for a type of PushAllowance node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#pushallowance.
	TypePushAllowance = "PushAllowance"

	// TypeReaction is constant for a type of Reaction node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reaction.
	TypeReaction = "Reaction"

	// TypeReadyForReviewEvent is constant for a type of ReadyForReviewEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#readyforreviewevent.
	TypeReadyForReviewEvent = "ReadyForReviewEvent"

	// TypeRef is constant for a type of Ref node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#ref.
	TypeRef = "Ref"

	// TypeReferencedEvent is constant for a type of ReferencedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#referencedevent.
	TypeReferencedEvent = "ReferencedEvent"

	// TypeRelease is constant for a type of Release node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#release.
	TypeRelease = "Release"

	// TypeReleaseAsset is constant for a type of ReleaseAsset node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#releaseasset.
	TypeReleaseAsset = "ReleaseAsset"

	// TypeRemovedFromProjectEvent is constant for a type of RemovedFromProjectEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#removedfromprojectevent.
	TypeRemovedFromProjectEvent = "RemovedFromProjectEvent"

	// TypeRenamedTitleEvent is constant for a type of RenamedTitleEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#renamedtitleevent.
	TypeRenamedTitleEvent = "RenamedTitleEvent"

	// TypeReopenedEvent is constant for a type of ReopenedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reopenedevent.
	TypeReopenedEvent = "ReopenedEvent"

	// TypeRepoAccessAuditEntry is constant for a type of RepoAccessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoaccessauditentry.
	TypeRepoAccessAuditEntry = "RepoAccessAuditEntry"

	// TypeRepoAddMemberAuditEntry is constant for a type of RepoAddMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoaddmemberauditentry.
	TypeRepoAddMemberAuditEntry = "RepoAddMemberAuditEntry"

	// TypeRepoAddTopicAuditEntry is constant for a type of RepoAddTopicAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoaddtopicauditentry.
	TypeRepoAddTopicAuditEntry = "RepoAddTopicAuditEntry"

	// TypeRepoArchivedAuditEntry is constant for a type of RepoArchivedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoarchivedauditentry.
	TypeRepoArchivedAuditEntry = "RepoArchivedAuditEntry"

	// TypeRepoChangeMergeSettingAuditEntry is constant for a type of RepoChangeMergeSettingAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repochangemergesettingauditentry.
	TypeRepoChangeMergeSettingAuditEntry = "RepoChangeMergeSettingAuditEntry"

	// TypeRepoConfigDisableAnonymousGitAccessAuditEntry is constant for a type of RepoConfigDisableAnonymousGitAccessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigdisableanonymousgitaccessauditentry.
	TypeRepoConfigDisableAnonymousGitAccessAuditEntry = "RepoConfigDisableAnonymousGitAccessAuditEntry"

	// TypeRepoConfigDisableCollaboratorsOnlyAuditEntry is constant for a type of RepoConfigDisableCollaboratorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigdisablecollaboratorsonlyauditentry.
	TypeRepoConfigDisableCollaboratorsOnlyAuditEntry = "RepoConfigDisableCollaboratorsOnlyAuditEntry"

	// TypeRepoConfigDisableContributorsOnlyAuditEntry is constant for a type of RepoConfigDisableContributorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigdisablecontributorsonlyauditentry.
	TypeRepoConfigDisableContributorsOnlyAuditEntry = "RepoConfigDisableContributorsOnlyAuditEntry"

	// TypeRepoConfigDisableSockpuppetDisallowedAuditEntry is constant for a type of RepoConfigDisableSockpuppetDisallowedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigdisablesockpuppetdisallowedauditentry.
	TypeRepoConfigDisableSockpuppetDisallowedAuditEntry = "RepoConfigDisableSockpuppetDisallowedAuditEntry"

	// TypeRepoConfigEnableAnonymousGitAccessAuditEntry is constant for a type of RepoConfigEnableAnonymousGitAccessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigenableanonymousgitaccessauditentry.
	TypeRepoConfigEnableAnonymousGitAccessAuditEntry = "RepoConfigEnableAnonymousGitAccessAuditEntry"

	// TypeRepoConfigEnableCollaboratorsOnlyAuditEntry is constant for a type of RepoConfigEnableCollaboratorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigenablecollaboratorsonlyauditentry.
	TypeRepoConfigEnableCollaboratorsOnlyAuditEntry = "RepoConfigEnableCollaboratorsOnlyAuditEntry"

	// TypeRepoConfigEnableContributorsOnlyAuditEntry is constant for a type of RepoConfigEnableContributorsOnlyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigenablecontributorsonlyauditentry.
	TypeRepoConfigEnableContributorsOnlyAuditEntry = "RepoConfigEnableContributorsOnlyAuditEntry"

	// TypeRepoConfigEnableSockpuppetDisallowedAuditEntry is constant for a type of RepoConfigEnableSockpuppetDisallowedAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigenablesockpuppetdisallowedauditentry.
	TypeRepoConfigEnableSockpuppetDisallowedAuditEntry = "RepoConfigEnableSockpuppetDisallowedAuditEntry"

	// TypeRepoConfigLockAnonymousGitAccessAuditEntry is constant for a type of RepoConfigLockAnonymousGitAccessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfiglockanonymousgitaccessauditentry.
	TypeRepoConfigLockAnonymousGitAccessAuditEntry = "RepoConfigLockAnonymousGitAccessAuditEntry"

	// TypeRepoConfigUnlockAnonymousGitAccessAuditEntry is constant for a type of RepoConfigUnlockAnonymousGitAccessAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repoconfigunlockanonymousgitaccessauditentry.
	TypeRepoConfigUnlockAnonymousGitAccessAuditEntry = "RepoConfigUnlockAnonymousGitAccessAuditEntry"

	// TypeRepoCreateAuditEntry is constant for a type of RepoCreateAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repocreateauditentry.
	TypeRepoCreateAuditEntry = "RepoCreateAuditEntry"

	// TypeRepoDestroyAuditEntry is constant for a type of RepoDestroyAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repodestroyauditentry.
	TypeRepoDestroyAuditEntry = "RepoDestroyAuditEntry"

	// TypeRepoRemoveMemberAuditEntry is constant for a type of RepoRemoveMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reporemovememberauditentry.
	TypeRepoRemoveMemberAuditEntry = "RepoRemoveMemberAuditEntry"

	// TypeRepoRemoveTopicAuditEntry is constant for a type of RepoRemoveTopicAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reporemovetopicauditentry.
	TypeRepoRemoveTopicAuditEntry = "RepoRemoveTopicAuditEntry"

	// TypeRepository is constant for a type of Repository node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repository.
	TypeRepository = "Repository"

	// TypeRepositoryInvitation is constant for a type of RepositoryInvitation node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositoryinvitation.
	TypeRepositoryInvitation = "RepositoryInvitation"

	// TypeRepositoryMigration is constant for a type of RepositoryMigration node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositorymigration.
	TypeRepositoryMigration = "RepositoryMigration"

	// TypeRepositoryTopic is constant for a type of RepositoryTopic node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositorytopic.
	TypeRepositoryTopic = "RepositoryTopic"

	// TypeRepositoryVisibilityChangeDisableAuditEntry is constant for a type of RepositoryVisibilityChangeDisableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositoryvisibilitychangedisableauditentry.
	TypeRepositoryVisibilityChangeDisableAuditEntry = "RepositoryVisibilityChangeDisableAuditEntry"

	// TypeRepositoryVisibilityChangeEnableAuditEntry is constant for a type of RepositoryVisibilityChangeEnableAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositoryvisibilitychangeenableauditentry.
	TypeRepositoryVisibilityChangeEnableAuditEntry = "RepositoryVisibilityChangeEnableAuditEntry"

	// TypeRepositoryVulnerabilityAlert is constant for a type of RepositoryVulnerabilityAlert node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#repositoryvulnerabilityalert.
	TypeRepositoryVulnerabilityAlert = "RepositoryVulnerabilityAlert"

	// TypeReviewDismissalAllowance is constant for a type of ReviewDismissalAllowance node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reviewdismissalallowance.
	TypeReviewDismissalAllowance = "ReviewDismissalAllowance"

	// TypeReviewDismissedEvent is constant for a type of ReviewDismissedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reviewdismissedevent.
	TypeReviewDismissedEvent = "ReviewDismissedEvent"

	// TypeReviewRequest is constant for a type of ReviewRequest node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reviewrequest.
	TypeReviewRequest = "ReviewRequest"

	// TypeReviewRequestRemovedEvent is constant for a type of ReviewRequestRemovedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reviewrequestremovedevent.
	TypeReviewRequestRemovedEvent = "ReviewRequestRemovedEvent"

	// TypeReviewRequestedEvent is constant for a type of ReviewRequestedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#reviewrequestedevent.
	TypeReviewRequestedEvent = "ReviewRequestedEvent"

	// TypeSavedReply is constant for a type of SavedReply node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#savedreply.
	TypeSavedReply = "SavedReply"

	// TypeSecurityAdvisory is constant for a type of SecurityAdvisory node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#securityadvisory.
	TypeSecurityAdvisory = "SecurityAdvisory"

	// TypeSponsorsActivity is constant for a type of SponsorsActivity node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorsactivity.
	TypeSponsorsActivity = "SponsorsActivity"

	// TypeSponsorsListing is constant for a type of SponsorsListing node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorslisting.
	TypeSponsorsListing = "SponsorsListing"

	// TypeSponsorsListingFeaturedItem is constant for a type of SponsorsListingFeaturedItem node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorslistingfeatureditem.
	TypeSponsorsListingFeaturedItem = "SponsorsListingFeaturedItem"

	// TypeSponsorsTier is constant for a type of SponsorsTier node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorstier.
	TypeSponsorsTier = "SponsorsTier"

	// TypeSponsorship is constant for a type of Sponsorship node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorship.
	TypeSponsorship = "Sponsorship"

	// TypeSponsorshipNewsletter is constant for a type of SponsorshipNewsletter node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#sponsorshipnewsletter.
	TypeSponsorshipNewsletter = "SponsorshipNewsletter"

	// TypeStatus is constant for a type of Status node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#status.
	TypeStatus = "Status"

	// TypeStatusCheckRollup is constant for a type of StatusCheckRollup node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#statuscheckrollup.
	TypeStatusCheckRollup = "StatusCheckRollup"

	// TypeStatusContext is constant for a type of StatusContext node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#statuscontext.
	TypeStatusContext = "StatusContext"

	// TypeSubscribedEvent is constant for a type of SubscribedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#subscribedevent.
	TypeSubscribedEvent = "SubscribedEvent"

	// TypeTag is constant for a type of Tag node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#tag.
	TypeTag = "Tag"

	// TypeTeam is constant for a type of Team node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#team.
	TypeTeam = "Team"

	// TypeTeamAddMemberAuditEntry is constant for a type of TeamAddMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamaddmemberauditentry.
	TypeTeamAddMemberAuditEntry = "TeamAddMemberAuditEntry"

	// TypeTeamAddRepositoryAuditEntry is constant for a type of TeamAddRepositoryAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamaddrepositoryauditentry.
	TypeTeamAddRepositoryAuditEntry = "TeamAddRepositoryAuditEntry"

	// TypeTeamChangeParentTeamAuditEntry is constant for a type of TeamChangeParentTeamAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamchangeparentteamauditentry.
	TypeTeamChangeParentTeamAuditEntry = "TeamChangeParentTeamAuditEntry"

	// TypeTeamDiscussion is constant for a type of TeamDiscussion node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamdiscussion.
	TypeTeamDiscussion = "TeamDiscussion"

	// TypeTeamDiscussionComment is constant for a type of TeamDiscussionComment node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamdiscussioncomment.
	TypeTeamDiscussionComment = "TeamDiscussionComment"

	// TypeTeamRemoveMemberAuditEntry is constant for a type of TeamRemoveMemberAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamremovememberauditentry.
	TypeTeamRemoveMemberAuditEntry = "TeamRemoveMemberAuditEntry"

	// TypeTeamRemoveRepositoryAuditEntry is constant for a type of TeamRemoveRepositoryAuditEntry node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#teamremoverepositoryauditentry.
	TypeTeamRemoveRepositoryAuditEntry = "TeamRemoveRepositoryAuditEntry"

	// TypeTopic is constant for a type of Topic node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#topic.
	TypeTopic = "Topic"

	// TypeTransferredEvent is constant for a type of TransferredEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#transferredevent.
	TypeTransferredEvent = "TransferredEvent"

	// TypeTree is constant for a type of Tree node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#tree.
	TypeTree = "Tree"

	// TypeUnassignedEvent is constant for a type of UnassignedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unassignedevent.
	TypeUnassignedEvent = "UnassignedEvent"

	// TypeUnlabeledEvent is constant for a type of UnlabeledEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unlabeledevent.
	TypeUnlabeledEvent = "UnlabeledEvent"

	// TypeUnlockedEvent is constant for a type of UnlockedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unlockedevent.
	TypeUnlockedEvent = "UnlockedEvent"

	// TypeUnmarkedAsDuplicateEvent is constant for a type of UnmarkedAsDuplicateEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unmarkedasduplicateevent.
	TypeUnmarkedAsDuplicateEvent = "UnmarkedAsDuplicateEvent"

	// TypeUnpinnedEvent is constant for a type of UnpinnedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unpinnedevent.
	TypeUnpinnedEvent = "UnpinnedEvent"

	// TypeUnsubscribedEvent is constant for a type of UnsubscribedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#unsubscribedevent.
	TypeUnsubscribedEvent = "UnsubscribedEvent"

	// TypeUser is constant for a type of User node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#user.
	TypeUser = "User"

	// TypeUserBlockedEvent is constant for a type of UserBlockedEvent node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#userblockedevent.
	TypeUserBlockedEvent = "UserBlockedEvent"

	// TypeUserContentEdit is constant for a type of UserContentEdit node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#usercontentedit.
	TypeUserContentEdit = "UserContentEdit"

	// TypeUserStatus is constant for a type of UserStatus node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#userstatus.
	TypeUserStatus = "UserStatus"

	// TypeVerifiableDomain is constant for a type of VerifiableDomain node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#verifiabledomain.
	TypeVerifiableDomain = "VerifiableDomain"

	// TypeWorkflow is constant for a type of Workflow node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#workflow.
	TypeWorkflow = "Workflow"

	// TypeWorkflowRun is constant for a type of WorkflowRun node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#workflowrun.
	TypeWorkflowRun = "WorkflowRun"
)
