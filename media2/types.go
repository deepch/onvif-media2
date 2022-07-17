package media2

import (
	"github.com/deepch/onvif-media2/xsd"
	"github.com/deepch/onvif-media2/xsd/onvif"
)

type Capabilities struct {
	SnapshotUri           bool `xml:"SnapshotUri,attr"`
	Rotation              bool `xml:"Rotation,attr"`
	VideoSourceMode       bool `xml:"VideoSourceMode,attr"`
	OSD                   bool `xml:"OSD,attr"`
	TemporaryOSDText      bool `xml:"TemporaryOSDText,attr"`
	EXICompression        bool `xml:"EXICompression,attr"`
	ProfileCapabilities   ProfileCapabilities
	StreamingCapabilities StreamingCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type StreamingCapabilities struct {
	RTPMulticast        bool `xml:"RTPMulticast,attr"`
	RTP_TCP             bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP        bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr"`
	NoRTSPStreaming     bool `xml:"NoRTSPStreaming,attr"`
}

//Media main types

type GetServiceCapabilities struct {
	XMLName string `xml:"tr2:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetVideoSources struct {
	XMLName string `xml:"tr2:GetVideoSources"`
}

type GetVideoSourcesResponse struct {
	VideoSources onvif.VideoSource
}

type GetAudioSources struct {
	XMLName string `xml:"tr2:GetAudioSources"`
}

type GetAudioSourcesResponse struct {
	AudioSources onvif.AudioSource
}

type GetAudioOutputs struct {
	XMLName string `xml:"tr2:GetAudioOutputs"`
}

type GetAudioOutputsResponse struct {
	AudioOutputs onvif.AudioOutput
}

type CreateProfile struct {
	XMLName string               `xml:"tr2:CreateProfile"`
	Name    onvif.Name           `xml:"tr2:Name"`
	Token   onvif.ReferenceToken `xml:"tr2:Token"`
}

type CreateProfileResponse struct {
	Profile onvif.Profile
}

type GetProfile struct {
	XMLName      string               `xml:"tr2:GetProfile"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetProfileResponse struct {
	Profile onvif.Profile
}

type GetProfiles struct {
	XMLName string `xml:"tr2:GetProfiles"`
}

type GetProfilesResponse struct {
	Profiles []onvif.Profile
}

type AddVideoEncoderConfiguration struct {
	XMLName            string               `xml:"tr2:AddVideoEncoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddVideoEncoderConfigurationResponse struct {
}

type RemoveVideoEncoderConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveVideoEncoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveVideoEncoderConfigurationResponse struct {
}

type AddVideoSourceConfiguration struct {
	XMLName            string               `xml:"tr2:AddVideoSourceConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddVideoSourceConfigurationResponse struct {
}

type RemoveVideoSourceConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveVideoSourceConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveVideoSourceConfigurationResponse struct {
}

type AddAudioEncoderConfiguration struct {
	XMLName            string               `xml:"tr2:AddAudioEncoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddAudioEncoderConfigurationResponse struct {
}

type RemoveAudioEncoderConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveAudioEncoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveAudioEncoderConfigurationResponse struct {
}

type AddAudioSourceConfiguration struct {
	XMLName            string               `xml:"tr2:AddAudioSourceConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddAudioSourceConfigurationResponse struct {
}

type RemoveAudioSourceConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveAudioSourceConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveAudioSourceConfigurationResponse struct {
}

type AddPTZConfiguration struct {
	XMLName            string               `xml:"tr2:AddPTZConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddPTZConfigurationResponse struct {
}

type RemovePTZConfiguration struct {
	XMLName      string               `xml:"tr2:RemovePTZConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemovePTZConfigurationResponse struct {
}

type AddVideoAnalyticsConfiguration struct {
	XMLName            string               `xml:"tr2:AddVideoAnalyticsConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddVideoAnalyticsConfigurationResponse struct {
}

type RemoveVideoAnalyticsConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveVideoAnalyticsConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveVideoAnalyticsConfigurationResponse struct {
}

type AddMetadataConfiguration struct {
	XMLName            string               `xml:"tr2:AddMetadataConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddMetadataConfigurationResponse struct {
}

type RemoveMetadataConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveMetadataConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveMetadataConfigurationResponse struct {
}

type AddAudioOutputConfiguration struct {
	XMLName            string               `xml:"tr2:AddAudioOutputConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddAudioOutputConfigurationResponse struct {
}

type RemoveAudioOutputConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveAudioOutputConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveAudioOutputConfigurationResponse struct {
}

type AddAudioDecoderConfiguration struct {
	XMLName            string               `xml:"tr2:AddAudioDecoderConfiguration"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type AddAudioDecoderConfigurationResponse struct {
}

type RemoveAudioDecoderConfiguration struct {
	XMLName      string               `xml:"tr2:RemoveAudioDecoderConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type RemoveAudioDecoderConfigurationResponse struct {
}

type DeleteProfile struct {
	XMLName      string               `xml:"tr2:DeleteProfile"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type DeleteProfileResponse struct {
}

type GetVideoSourceConfigurations struct {
	XMLName string `xml:"tr2:GetVideoSourceConfigurations"`
}

type GetVideoSourceConfigurationsResponse struct {
	Configurations onvif.VideoSourceConfiguration
}

type GetVideoEncoderConfigurations struct {
	XMLName string `xml:"tr2:GetVideoEncoderConfigurations"`
}

type GetVideoEncoderConfigurationsResponse struct {
	Configurations onvif.VideoEncoderConfiguration
}

type GetAudioSourceConfigurations struct {
	XMLName string `xml:"tr2:GetAudioSourceConfigurations"`
}

type GetAudioSourceConfigurationsResponse struct {
	Configurations onvif.AudioSourceConfiguration
}

type GetAudioEncoderConfigurations struct {
	XMLName string `xml:"tr2:GetAudioEncoderConfigurations"`
}

type GetAudioEncoderConfigurationsResponse struct {
	Configurations onvif.AudioEncoderConfiguration
}

type GetVideoAnalyticsConfigurations struct {
	XMLName string `xml:"tr2:GetVideoAnalyticsConfigurations"`
}

type GetVideoAnalyticsConfigurationsResponse struct {
	Configurations onvif.VideoAnalyticsConfiguration
}

type GetMetadataConfigurations struct {
	XMLName string `xml:"tr2:GetMetadataConfigurations"`
}

type GetMetadataConfigurationsResponse struct {
	Configurations onvif.MetadataConfiguration
}

type GetAudioOutputConfigurations struct {
	XMLName string `xml:"tr2:GetAudioOutputConfigurations"`
}

type GetAudioOutputConfigurationsResponse struct {
	Configurations onvif.AudioOutputConfiguration
}

type GetAudioDecoderConfigurations struct {
	XMLName string `xml:"tr2:GetAudioDecoderConfigurations"`
}

type GetAudioDecoderConfigurationsResponse struct {
	Configurations onvif.AudioDecoderConfiguration
}

type GetVideoSourceConfiguration struct {
	XMLName            string               `xml:"tr2:GetVideoSourceConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetVideoSourceConfigurationResponse struct {
	Configuration onvif.VideoSourceConfiguration
}

type GetVideoEncoderConfiguration struct {
	XMLName            string               `xml:"tr2:GetVideoEncoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetVideoEncoderConfigurationResponse struct {
	Configuration onvif.VideoEncoderConfiguration
}

type GetAudioSourceConfiguration struct {
	XMLName            string               `xml:"tr2:GetAudioSourceConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioSourceConfigurationResponse struct {
	Configuration onvif.AudioSourceConfiguration
}

type GetAudioEncoderConfiguration struct {
	XMLName            string               `xml:"tr2:GetAudioEncoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioEncoderConfigurationResponse struct {
	Configuration onvif.AudioEncoderConfiguration
}

type GetVideoAnalyticsConfiguration struct {
	XMLName            string               `xml:"tr2:GetVideoAnalyticsConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetVideoAnalyticsConfigurationResponse struct {
	Configuration onvif.VideoAnalyticsConfiguration
}

type GetMetadataConfiguration struct {
	XMLName            string               `xml:"tr2:GetMetadataConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetMetadataConfigurationResponse struct {
	Configuration onvif.MetadataConfiguration
}

type GetAudioOutputConfiguration struct {
	XMLName            string               `xml:"tr2:GetAudioOutputConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioOutputConfigurationResponse struct {
	Configuration onvif.AudioOutputConfiguration
}

type GetAudioDecoderConfiguration struct {
	XMLName            string               `xml:"tr2:GetAudioDecoderConfiguration"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioDecoderConfigurationResponse struct {
	Configuration onvif.AudioDecoderConfiguration
}

type GetCompatibleVideoEncoderConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleVideoEncoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleVideoEncoderConfigurationsResponse struct {
	Configurations onvif.VideoEncoderConfiguration
}

type GetCompatibleVideoSourceConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleVideoSourceConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleVideoSourceConfigurationsResponse struct {
	Configurations onvif.VideoSourceConfiguration
}

type GetCompatibleAudioEncoderConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleAudioEncoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleAudioEncoderConfigurationsResponse struct {
	Configurations onvif.AudioEncoderConfiguration
}

type GetCompatibleAudioSourceConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleAudioSourceConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleAudioSourceConfigurationsResponse struct {
	Configurations onvif.AudioSourceConfiguration
}

type GetCompatibleVideoAnalyticsConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleVideoAnalyticsConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	Configurations onvif.VideoAnalyticsConfiguration
}

type GetCompatibleMetadataConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleMetadataConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleMetadataConfigurationsResponse struct {
	Configurations onvif.MetadataConfiguration
}

type GetCompatibleAudioOutputConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleAudioOutputConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleAudioOutputConfigurationsResponse struct {
	Configurations onvif.AudioOutputConfiguration
}

type GetCompatibleAudioDecoderConfigurations struct {
	XMLName      string               `xml:"tr2:GetCompatibleAudioDecoderConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetCompatibleAudioDecoderConfigurationsResponse struct {
	Configurations onvif.AudioDecoderConfiguration
}

type SetVideoSourceConfiguration struct {
	XMLName          string                         `xml:"tr2:SetVideoSourceConfiguration"`
	Configuration    onvif.VideoSourceConfiguration `xml:"tr2:Configuration"`
	ForcePersistence xsd.Boolean                    `xml:"tr2:ForcePersistence"`
}

type SetVideoSourceConfigurationResponse struct {
}

type SetVideoEncoderConfiguration struct {
	XMLName       string                           `xml:"tr2:SetVideoEncoderConfiguration"`
	Configuration onvif.VideoEncoder2Configuration `xml:"tr2:Configuration"`
	//ForcePersistence xsd.Boolean                     `xml:"tr2:ForcePersistence"`
}

type SetVideoEncoderConfigurationResponse struct {
}

type SetAudioSourceConfiguration struct {
	XMLName          string                         `xml:"tr2:SetAudioSourceConfiguration"`
	Configuration    onvif.AudioSourceConfiguration `xml:"tr2:Configuration"`
	ForcePersistence xsd.Boolean                    `xml:"tr2:ForcePersistence"`
}

type SetAudioSourceConfigurationResponse struct {
}

type SetAudioEncoderConfiguration struct {
	XMLName          string                          `xml:"tr2:SetAudioEncoderConfiguration"`
	Configuration    onvif.AudioEncoderConfiguration `xml:"tr2:Configuration"`
	ForcePersistence xsd.Boolean                     `xml:"tr2:ForcePersistence"`
}

type SetAudioEncoderConfigurationResponse struct {
}

type SetVideoAnalyticsConfiguration struct {
	XMLName          string                            `xml:"tr2:SetVideoAnalyticsConfiguration"`
	Configuration    onvif.VideoAnalyticsConfiguration `xml:"tr2:Configuration"`
	ForcePersistence bool                              `xml:"tr2:ForcePersistence"`
}

type SetVideoAnalyticsConfigurationResponse struct {
}

type SetMetadataConfiguration struct {
	XMLName          string                      `xml:"tr2:GetDeviceInformation"`
	Configuration    onvif.MetadataConfiguration `xml:"tr2:Configuration"`
	ForcePersistence xsd.Boolean                 `xml:"tr2:ForcePersistence"`
}

type SetMetadataConfigurationResponse struct {
}

type SetAudioOutputConfiguration struct {
	XMLName          string                         `xml:"tr2:SetAudioOutputConfiguration"`
	Configuration    onvif.AudioOutputConfiguration `xml:"tr2:Configuration"`
	ForcePersistence bool                           `xml:"tr2:ForcePersistence"`
}

type SetAudioOutputConfigurationResponse struct {
}

type SetAudioDecoderConfiguration struct {
	XMLName          string                          `xml:"tr2:SetAudioDecoderConfiguration"`
	Configuration    onvif.AudioDecoderConfiguration `xml:"tr2:Configuration"`
	ForcePersistence xsd.Boolean                     `xml:"tr2:ForcePersistence"`
}

type SetAudioDecoderConfigurationResponse struct {
}

type GetVideoSourceConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetVideoSourceConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetVideoSourceConfigurationOptionsResponse struct {
	Options onvif.VideoSourceConfigurationOptions
}

type GetVideoEncoderConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetVideoEncoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetVideoEncoderConfigurationOptionsResponse struct {
	Options onvif.VideoEncoderConfigurationOptions
}

type GetAudioSourceConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetAudioSourceConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioSourceConfigurationOptionsResponse struct {
	Options onvif.AudioSourceConfigurationOptions
}

type GetAudioEncoderConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetAudioEncoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioEncoderConfigurationOptionsResponse struct {
	Options onvif.AudioEncoderConfigurationOptions
}

type GetMetadataConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetMetadataConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetMetadataConfigurationOptionsResponse struct {
	Options onvif.MetadataConfigurationOptions
}

type GetAudioOutputConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetAudioOutputConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioOutputConfigurationOptionsResponse struct {
	Options onvif.AudioOutputConfigurationOptions
}

type GetAudioDecoderConfigurationOptions struct {
	XMLName            string               `xml:"tr2:GetAudioDecoderConfigurationOptions"`
	ProfileToken       onvif.ReferenceToken `xml:"tr2:ProfileToken"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetAudioDecoderConfigurationOptionsResponse struct {
	Options onvif.AudioDecoderConfigurationOptions
}

type GetGuaranteedNumberOfVideoEncoderInstances struct {
	XMLName            string               `xml:"tr2:GetGuaranteedNumberOfVideoEncoderInstances"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	TotalNumber int
	JPEG        int
	H264        int
	MPEG4       int
}

type GetStreamUri struct {
	XMLName      string               `xml:"tr2:GetStreamUri"`
	StreamSetup  onvif.StreamSetup    `xml:"tr2:StreamSetup"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetStreamUriResponse struct {
	MediaUri onvif.MediaUri
}

type StartMulticastStreaming struct {
	XMLName      string               `xml:"tr2:StartMulticastStreaming"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type StartMulticastStreamingResponse struct {
}

type StopMulticastStreaming struct {
	XMLName      string               `xml:"tr2:StopMulticastStreaming"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type StopMulticastStreamingResponse struct {
}

type SetSynchronizationPoint struct {
	XMLName      string               `xml:"tr2:SetSynchronizationPoint"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type SetSynchronizationPointResponse struct {
}

type GetSnapshotUri struct {
	XMLName      string               `xml:"tr2:GetSnapshotUri"`
	ProfileToken onvif.ReferenceToken `xml:"tr2:ProfileToken"`
}

type GetSnapshotUriResponse struct {
	MediaUri onvif.MediaUri
}

type GetVideoSourceModes struct {
	XMLName          string               `xml:"tr2:GetVideoSourceModes"`
	VideoSourceToken onvif.ReferenceToken `xml:"tr2:VideoSourceToken"`
}

type GetVideoSourceModesResponse struct {
	VideoSourceModes onvif.VideoSourceMode
}

type SetVideoSourceMode struct {
	XMLName              string               `xml:"tr2:SetVideoSourceMode"`
	VideoSourceToken     onvif.ReferenceToken `xml:"tr2:VideoSourceToken"`
	VideoSourceModeToken onvif.ReferenceToken `xml:"tr2:VideoSourceModeToken"`
}

type SetVideoSourceModeResponse struct {
	Reboot bool
}

type GetOSDs struct {
	XMLName            string               `xml:"tr2:GetOSDs"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetOSDsResponse struct {
	OSDs onvif.OSDConfiguration
}

type GetOSD struct {
	XMLName  string               `xml:"tr2:GetOSD"`
	OSDToken onvif.ReferenceToken `xml:"tr2:OSDToken"`
}

type GetOSDResponse struct {
	OSD onvif.OSDConfiguration
}

type GetOSDOptions struct {
	XMLName            string               `xml:"tr2:GetOSDOptions"`
	ConfigurationToken onvif.ReferenceToken `xml:"tr2:ConfigurationToken"`
}

type GetOSDOptionsResponse struct {
	OSDOptions onvif.OSDConfigurationOptions
}

type SetOSD struct {
	XMLName string                 `xml:"tr2:SetOSD"`
	OSD     onvif.OSDConfiguration `xml:"tr2:OSD"`
}

type SetOSDResponse struct {
}

type CreateOSD struct {
	XMLName string                 `xml:"tr2:CreateOSD"`
	OSD     onvif.OSDConfiguration `xml:"tr2:OSD"`
}

type CreateOSDResponse struct {
	OSDToken onvif.ReferenceToken
}

type DeleteOSD struct {
	XMLName  string               `xml:"tr2:DeleteOSD"`
	OSDToken onvif.ReferenceToken `xml:"tr2:OSDToken"`
}

type DeleteOSDResponse struct {
}
