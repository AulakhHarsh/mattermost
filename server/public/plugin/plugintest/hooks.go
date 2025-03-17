// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make plugin-mocks`.

package plugintest

import (
	io "io"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"

	plugin "github.com/mattermost/mattermost/server/public/plugin"

	saml2 "github.com/mattermost/gosaml2"
)

// Hooks is an autogenerated mock type for the Hooks type
type Hooks struct {
	mock.Mock
}

// ChannelHasBeenCreated provides a mock function with given fields: c, channel
func (_m *Hooks) ChannelHasBeenCreated(c *plugin.Context, channel *model.Channel) {
	_m.Called(c, channel)
}

// ConfigurationWillBeSaved provides a mock function with given fields: newCfg
func (_m *Hooks) ConfigurationWillBeSaved(newCfg *model.Config) (*model.Config, error) {
	ret := _m.Called(newCfg)

	if len(ret) == 0 {
		panic("no return value specified for ConfigurationWillBeSaved")
	}

	var r0 *model.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Config) (*model.Config, error)); ok {
		return rf(newCfg)
	}
	if rf, ok := ret.Get(0).(func(*model.Config) *model.Config); ok {
		r0 = rf(newCfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Config) error); ok {
		r1 = rf(newCfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteCommand provides a mock function with given fields: c, args
func (_m *Hooks) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	ret := _m.Called(c, args)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteCommand")
	}

	var r0 *model.CommandResponse
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.CommandArgs) (*model.CommandResponse, *model.AppError)); ok {
		return rf(c, args)
	}
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.CommandArgs) *model.CommandResponse); ok {
		r0 = rf(c, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CommandResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.CommandArgs) *model.AppError); ok {
		r1 = rf(c, args)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// FileWillBeUploaded provides a mock function with given fields: c, info, file, output
func (_m *Hooks) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	ret := _m.Called(c, info, file, output)

	if len(ret) == 0 {
		panic("no return value specified for FileWillBeUploaded")
	}

	var r0 *model.FileInfo
	var r1 string
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.FileInfo, io.Reader, io.Writer) (*model.FileInfo, string)); ok {
		return rf(c, info, file, output)
	}
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.FileInfo, io.Reader, io.Writer) *model.FileInfo); ok {
		r0 = rf(c, info, file, output)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.FileInfo, io.Reader, io.Writer) string); ok {
		r1 = rf(c, info, file, output)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GenerateSupportData provides a mock function with given fields: c
func (_m *Hooks) GenerateSupportData(c *plugin.Context) ([]*model.FileData, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GenerateSupportData")
	}

	var r0 []*model.FileData
	var r1 error
	if rf, ok := ret.Get(0).(func(*plugin.Context) ([]*model.FileData, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(*plugin.Context) []*model.FileData); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FileData)
		}
	}

	if rf, ok := ret.Get(1).(func(*plugin.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Implemented provides a mock function with given fields:
func (_m *Hooks) Implemented() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Implemented")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageHasBeenDeleted provides a mock function with given fields: c, post
func (_m *Hooks) MessageHasBeenDeleted(c *plugin.Context, post *model.Post) {
	_m.Called(c, post)
}

// MessageHasBeenPosted provides a mock function with given fields: c, post
func (_m *Hooks) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	_m.Called(c, post)
}

// MessageHasBeenUpdated provides a mock function with given fields: c, newPost, oldPost
func (_m *Hooks) MessageHasBeenUpdated(c *plugin.Context, newPost *model.Post, oldPost *model.Post) {
	_m.Called(c, newPost, oldPost)
}

// MessageWillBePosted provides a mock function with given fields: c, post
func (_m *Hooks) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	ret := _m.Called(c, post)

	if len(ret) == 0 {
		panic("no return value specified for MessageWillBePosted")
	}

	var r0 *model.Post
	var r1 string
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post) (*model.Post, string)); ok {
		return rf(c, post)
	}
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post) *model.Post); ok {
		r0 = rf(c, post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.Post) string); ok {
		r1 = rf(c, post)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MessageWillBeUpdated provides a mock function with given fields: c, newPost, oldPost
func (_m *Hooks) MessageWillBeUpdated(c *plugin.Context, newPost *model.Post, oldPost *model.Post) (*model.Post, string) {
	ret := _m.Called(c, newPost, oldPost)

	if len(ret) == 0 {
		panic("no return value specified for MessageWillBeUpdated")
	}

	var r0 *model.Post
	var r1 string
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post, *model.Post) (*model.Post, string)); ok {
		return rf(c, newPost, oldPost)
	}
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.Post, *model.Post) *model.Post); ok {
		r0 = rf(c, newPost, oldPost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(*plugin.Context, *model.Post, *model.Post) string); ok {
		r1 = rf(c, newPost, oldPost)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MessagesWillBeConsumed provides a mock function with given fields: posts
func (_m *Hooks) MessagesWillBeConsumed(posts []*model.Post) []*model.Post {
	ret := _m.Called(posts)

	if len(ret) == 0 {
		panic("no return value specified for MessagesWillBeConsumed")
	}

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func([]*model.Post) []*model.Post); ok {
		r0 = rf(posts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	return r0
}

// NotificationWillBePushed provides a mock function with given fields: pushNotification, userID
func (_m *Hooks) NotificationWillBePushed(pushNotification *model.PushNotification, userID string) (*model.PushNotification, string) {
	ret := _m.Called(pushNotification, userID)

	if len(ret) == 0 {
		panic("no return value specified for NotificationWillBePushed")
	}

	var r0 *model.PushNotification
	var r1 string
	if rf, ok := ret.Get(0).(func(*model.PushNotification, string) (*model.PushNotification, string)); ok {
		return rf(pushNotification, userID)
	}
	if rf, ok := ret.Get(0).(func(*model.PushNotification, string) *model.PushNotification); ok {
		r0 = rf(pushNotification, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PushNotification)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.PushNotification, string) string); ok {
		r1 = rf(pushNotification, userID)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// OnActivate provides a mock function with given fields:
func (_m *Hooks) OnActivate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnActivate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnCloudLimitsUpdated provides a mock function with given fields: limits
func (_m *Hooks) OnCloudLimitsUpdated(limits *model.ProductLimits) {
	_m.Called(limits)
}

// OnConfigurationChange provides a mock function with given fields:
func (_m *Hooks) OnConfigurationChange() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnConfigurationChange")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnDeactivate provides a mock function with given fields:
func (_m *Hooks) OnDeactivate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnDeactivate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnInstall provides a mock function with given fields: c, event
func (_m *Hooks) OnInstall(c *plugin.Context, event model.OnInstallEvent) error {
	ret := _m.Called(c, event)

	if len(ret) == 0 {
		panic("no return value specified for OnInstall")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*plugin.Context, model.OnInstallEvent) error); ok {
		r0 = rf(c, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnPluginClusterEvent provides a mock function with given fields: c, ev
func (_m *Hooks) OnPluginClusterEvent(c *plugin.Context, ev model.PluginClusterEvent) {
	_m.Called(c, ev)
}

// OnSAMLLogin provides a mock function with given fields: c, user, assertion
func (_m *Hooks) OnSAMLLogin(c *plugin.Context, user *model.User, assertion *saml2.AssertionInfo) error {
	ret := _m.Called(c, user, assertion)

	if len(ret) == 0 {
		panic("no return value specified for OnSAMLLogin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.User, *saml2.AssertionInfo) error); ok {
		r0 = rf(c, user, assertion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnSendDailyTelemetry provides a mock function with given fields:
func (_m *Hooks) OnSendDailyTelemetry() {
	_m.Called()
}

// OnSharedChannelsAttachmentSyncMsg provides a mock function with given fields: fi, post, rc
func (_m *Hooks) OnSharedChannelsAttachmentSyncMsg(fi *model.FileInfo, post *model.Post, rc *model.RemoteCluster) error {
	ret := _m.Called(fi, post, rc)

	if len(ret) == 0 {
		panic("no return value specified for OnSharedChannelsAttachmentSyncMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.FileInfo, *model.Post, *model.RemoteCluster) error); ok {
		r0 = rf(fi, post, rc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnSharedChannelsPing provides a mock function with given fields: rc
func (_m *Hooks) OnSharedChannelsPing(rc *model.RemoteCluster) bool {
	ret := _m.Called(rc)

	if len(ret) == 0 {
		panic("no return value specified for OnSharedChannelsPing")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*model.RemoteCluster) bool); ok {
		r0 = rf(rc)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OnSharedChannelsProfileImageSyncMsg provides a mock function with given fields: user, rc
func (_m *Hooks) OnSharedChannelsProfileImageSyncMsg(user *model.User, rc *model.RemoteCluster) error {
	ret := _m.Called(user, rc)

	if len(ret) == 0 {
		panic("no return value specified for OnSharedChannelsProfileImageSyncMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.RemoteCluster) error); ok {
		r0 = rf(user, rc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnSharedChannelsSyncMsg provides a mock function with given fields: msg, rc
func (_m *Hooks) OnSharedChannelsSyncMsg(msg *model.SyncMsg, rc *model.RemoteCluster) (model.SyncResponse, error) {
	ret := _m.Called(msg, rc)

	if len(ret) == 0 {
		panic("no return value specified for OnSharedChannelsSyncMsg")
	}

	var r0 model.SyncResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.SyncMsg, *model.RemoteCluster) (model.SyncResponse, error)); ok {
		return rf(msg, rc)
	}
	if rf, ok := ret.Get(0).(func(*model.SyncMsg, *model.RemoteCluster) model.SyncResponse); ok {
		r0 = rf(msg, rc)
	} else {
		r0 = ret.Get(0).(model.SyncResponse)
	}

	if rf, ok := ret.Get(1).(func(*model.SyncMsg, *model.RemoteCluster) error); ok {
		r1 = rf(msg, rc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnWebSocketConnect provides a mock function with given fields: webConnID, userID
func (_m *Hooks) OnWebSocketConnect(webConnID string, userID string) {
	_m.Called(webConnID, userID)
}

// OnWebSocketDisconnect provides a mock function with given fields: webConnID, userID
func (_m *Hooks) OnWebSocketDisconnect(webConnID string, userID string) {
	_m.Called(webConnID, userID)
}

// PreferencesHaveChanged provides a mock function with given fields: c, preferences
func (_m *Hooks) PreferencesHaveChanged(c *plugin.Context, preferences []model.Preference) {
	_m.Called(c, preferences)
}

// ReactionHasBeenAdded provides a mock function with given fields: c, reaction
func (_m *Hooks) ReactionHasBeenAdded(c *plugin.Context, reaction *model.Reaction) {
	_m.Called(c, reaction)
}

// ReactionHasBeenRemoved provides a mock function with given fields: c, reaction
func (_m *Hooks) ReactionHasBeenRemoved(c *plugin.Context, reaction *model.Reaction) {
	_m.Called(c, reaction)
}

// RunDataRetention provides a mock function with given fields: nowTime, batchSize
func (_m *Hooks) RunDataRetention(nowTime int64, batchSize int64) (int64, error) {
	ret := _m.Called(nowTime, batchSize)

	if len(ret) == 0 {
		panic("no return value specified for RunDataRetention")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (int64, error)); ok {
		return rf(nowTime, batchSize)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(nowTime, batchSize)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(nowTime, batchSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServeHTTP provides a mock function with given fields: c, w, r
func (_m *Hooks) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	_m.Called(c, w, r)
}

// ServeMetrics provides a mock function with given fields: c, w, r
func (_m *Hooks) ServeMetrics(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	_m.Called(c, w, r)
}

// UserHasBeenCreated provides a mock function with given fields: c, user
func (_m *Hooks) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	_m.Called(c, user)
}

// UserHasBeenDeactivated provides a mock function with given fields: c, user
func (_m *Hooks) UserHasBeenDeactivated(c *plugin.Context, user *model.User) {
	_m.Called(c, user)
}

// UserHasJoinedChannel provides a mock function with given fields: c, channelMember, actor
func (_m *Hooks) UserHasJoinedChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	_m.Called(c, channelMember, actor)
}

// UserHasJoinedTeam provides a mock function with given fields: c, teamMember, actor
func (_m *Hooks) UserHasJoinedTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	_m.Called(c, teamMember, actor)
}

// UserHasLeftChannel provides a mock function with given fields: c, channelMember, actor
func (_m *Hooks) UserHasLeftChannel(c *plugin.Context, channelMember *model.ChannelMember, actor *model.User) {
	_m.Called(c, channelMember, actor)
}

// UserHasLeftTeam provides a mock function with given fields: c, teamMember, actor
func (_m *Hooks) UserHasLeftTeam(c *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	_m.Called(c, teamMember, actor)
}

// UserHasLoggedIn provides a mock function with given fields: c, user
func (_m *Hooks) UserHasLoggedIn(c *plugin.Context, user *model.User) {
	_m.Called(c, user)
}

// UserWillLogIn provides a mock function with given fields: c, user
func (_m *Hooks) UserWillLogIn(c *plugin.Context, user *model.User) string {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for UserWillLogIn")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(*plugin.Context, *model.User) string); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WebSocketMessageHasBeenPosted provides a mock function with given fields: webConnID, userID, req
func (_m *Hooks) WebSocketMessageHasBeenPosted(webConnID string, userID string, req *model.WebSocketRequest) {
	_m.Called(webConnID, userID, req)
}

// NewHooks creates a new instance of Hooks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHooks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Hooks {
	mock := &Hooks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
