package xapi

import (
	"time"
)

type SR struct {
	//NameDescription string        `json:"name_description"`

	LocalCacheEnabled   bool     `xmlrpc:"local_cache_enabled"`
	NameLabel           string   `xmlrpc:"name_label" json:"name_label`
	PhysicalSize        string   `xmlrpc:"physical_size"`
	PhysicalUtilization string   `xmlrpc:"physical_utilisation"`
	Shared              bool     `xmlrpc:"shared"`
	Tags                []string `xmlrpc:"tags"`
	Type                string   `xmlrpc:"type"`
	UUID                string   `xmlrpc:"uuid" json:"uuid"`
	VirtualAllocation   string   `xmlrpc:"virtual_allocation"`
	VDIs                []string
	PDBs                []string
}

/*
	All commented fields are from an older version
	Currently this will only support XenServer 6.5
*/

type Session struct {
	AuthUserName     string            `xmlrpc:"auth_user_name"`
	AuthUserSid      string            `xmlrpc:"auth_user_sid"`
	IsLocalSuperuser bool              `xmlrpc:"is_local_superuser"`
	OtherConfig      map[string]string `xmlrpc:"other_config"`
	Parent           string
	Pool             bool
	RbacPermissions  []string `xmlrpc:"rbac_permissions"`
	Subject          string
	Tasks            []string
	ThisHost         string `xmlrpc:"this_host"`
	ThisUser         string `xmlrpc:"this_user"`
	UUID             string `xmlrpc:"uuid"`
	//ValidationTime   time.Time `xmlrpc:"validation_time"`
	//LastActive       time.Time         `xmlrpc:"last_active"`
}

type VDI struct {
	AllowCaching        bool              `xmlrpc:"Allow_caching"`
	AllowedOperations   []string          `xmlrpc:"Allowed_operations"`
	CrashDumps          []string          `xmlrpc:"Crash_dumps"`
	CurrentOperations   map[string]string `xmlrpc:"Current_operations"`
	IsAsnapshot         bool              `xmlrpc:"Is_a_snapshot"`
	Location            string
	Managed             bool
	MetadataLatest      bool `xmlrpc:"Metadata_latest"`
	Missing             bool
	NameDescription     string            `xmlrpc:"Name_description"`
	NameLabel           string            `xmlrpc:"Name_label"`
	OnBoot              string            `xmlrpc:"On_boot"`
	OtherConfig         map[string]string `xmlrpc:"Other_config"`
	Parent              string
	PhyscialUtilisation int  `xmlrpc:"Physcial_utilisation"`
	ReadOnly            bool `xmlrpc:"Read_only"`
	Sharable            bool
	SmConfig            map[string]string `xmlrpc:"Sm_config"`
	SnapshotOf          []string          `xmlrpc:"Snapshot_of"`
	SR                  []string
	StorageLock         bool `xmlrpc:"Storage_lock"`
	Tags                string
	Type                string `json:"type"`
	UUID                string `json:"uuid"`
	VDBs                []string
	VirtualSize         int               `xmlrpc:"Virtual_size"`
	XenstoreData        map[string]string `xmlrpc:"Xenstore_data"`
}

type VDB struct {
	AllowedOperations      []string          `xmlrpc:"Allowed_operations"`
	CurrentOperations      map[string]string `xmlrpc:"Current_operations"`
	Bootable               bool
	CurrentlyAttached      bool `xmlrpc:"Currently_attached"`
	Device                 string
	Empty                  bool
	Metrics                string
	Mode                   string
	OtherConfig            map[string]string `xmlrpc:"Other_config"`
	QosAlgorithmParams     map[string]string `xmlrpc:"Qos_algorithm_params"`
	QosAlgorithmType       string            `xmlrpc:"Qos_algorithm_type"`
	QosSupportedAlgorithms []string          `xmlrpc:"Qos_supported_algorithms"`
	RuntimeProperties      map[string]string `xmlrpc:"Runtime_properties"`
	StatusCode             int               `xmlrpc:"Status_code"`
	StatusDetail           string            `xmlrpc:"Status_detail"`
	UUID                   string            `json:"uuid"`
	VDI                    string
	VM                     string
	StorageLock            bool   `xmlrpc:"Storage_lock"`
	Type                   string `json:"type"`
	Unpluggable            bool
	Userdevice             string
}

/* Support for v1.3 */
type VM struct {
	ActionsAfterCrash    string `xmlrpc:"actions_after_crash"`
	ActionsAfterReboot   string `xmlrpc:"actions_after_reboot"`
	ActionsAfterShutdown string `xmlrpc:"actions_after_shutdown"`
	//Affinity                string
	AllowedOperations []string `xmlrpc:"allowed_operations"`

	Appliance    string   `xmlrpc:"appliance"`
	AttachedPCIs []string `xmlrpc:"attached_PCIs"`
	//BiosStrings  map[string]string `xmlrpc:"bios_strings"`
	/*
		Blobs                   map[string][]byte
		BlockedOperations       map[string]string `xmlrpc:"Blocked_operations"`
		Children                []string
		Consoles                []string
		CrashDumps              []string          `xmlrpc:"Crash_dumps"`
		CurrentOperations       map[string]string `xmlrpc:"Current_operations"`
		Domarch                 string
		Domid                   int
		GuestMetrics            string            `xmlrpc:"Guest_metrics"`
	*/
	HaAlwaysRun bool `xmlrpc:"ha_always_run"`
	//HaRestartPriority string `xmlrpc:"ha_restart_priority"`
	/*
		HVMBootParams           map[string]string `xmlrpc:"HVM_boot_params"`
		HVMBootPolicy           string            `xmlrpc:"HVM_boot_policy"`
	*/
	HVMShadowMultiplier float64           `xmlrpc:"HVM_shadow_multiplier"`
	IsASnapshot         bool              `xmlrpc:"is_a_snapshot"`
	IsATemplate         bool              `xmlrpc:"is_a_template"`
	IsControlDomain     bool              `xmlrpc:"is_control_domain"`
	IsSnapshotFromVmpp  bool              `xmlrpc:"is_snapshot_from_vmpp"`
	LastBootCPUFlags    map[string]string `xmlrpc:"last_boot_CPU_flags"`
	/*
		LastBootedRecord        string            `xmlrpc:"Last_booted_record"`
		MemoryDynamicMax        int               `xmlrpc:"Memory_dynamic_max"`
		MemoryDynamicMin        int               `xmlrpc:"Memory_dynamic_min"`
		MemoryOverhead          int               `xmlrpc:"Memory_overhead"`
		MemoryStaticMax         int               `xmlrpc:"Memory_static_max"`
		MemoryStaticMin         int               `xmlrpc:"Memory_static_min"`
		MemoryTarget            int               `xmlrpc:"Memory_target"`
	*/
	//Metrics         string `xmlrpc:"metrics"`
	//NameDescription string `xmlrpc:"name_description"`
	NameLabel string `xmlrpc:"name_label" json:"name_label`
	/*
		Order                   int
	*/
	//OtherConfig map[string]string `xmlrpc:"other_config"`
	/*
		Parent                  string
		PCIBus                  string `xmlrpc:"PCI_bus"`
		Platform                map[string]string
	*/
	PowerState string `xmlrpc:"power_state" json:"power_state"`
	//ProtectionPolicy        string `xmlrpc:"Protection_policy"`
	//PVArgs       string `xmlrpc:"PV_args"`
	//PVBootloader string `xmlrpc:"PV_bootloader"`
	//PVBootloaderArgs string `xmlrpc:"PV_bootloader_args"`
	//PVKernel string `xmlrpc:"PV_kernel"`
	/*
		PVLegacyArgs            string `xmlrpc:"PV_legacy_args"`
		PVRamdisk               string `xmlrpc:"PV_ramdisk"`
		Recommendations         string
		ResidentOn              string            `xmlrpc:"Resident_on"`
		ShutdownDelay           int               `xmlrpc:"shutdown_delay"`
		SnapshotInfo            map[string]string `xmlrpc:"Snapshot_info"`
	*/
	//SnapshotMetadata string `xmlrpc:"snapshot_metadata"`
	//SnapshotTime            time.Time `xmlrpc:"snapshot_time"`
	//StartDelay              int      `xmlrpc:"start_delay"`
	//SuspendSR               string   `xmlrpc:"suspend_SR"`
	//SuspendVDI              string   `xmlrpc:"suspend_VDI"`
	Tags []string `xmlrpc:"tags"`
	//TransportableSnapshotID string   `xmlrpc:"transportable_snapshot_id"`
	UserVersion string   `xmlrpc:"user_version"`
	UUID        string   `xmlrpc:"uuid" json:"uuid"`
	Version     string   `xmlrpc:"version"`
	VBDs        []string `xmlrpc:"VBDs"`
	/*
		VCPUsAtStartup          int               `xmlrpc:"VCPUs_at_startup"`
		VCPUsMax                int               `xmlrpc:"VCPUs_max"`
		VCPUsParams             map[string]string `xmlrpc:"VCPUs_params"`
		VGPUs                   []string
		VIFs                    []string
		VTPM                    []string
	*/
	//XenstoreData map[string]string `xmlrpc:"xenstore_data"`
}

type Event struct {
	Class     string
	ID        int    `xmlrpc:"id"`
	ObjUUID   string `xmlrpc:"Obj_uuid"`
	Operation string
	Ref       string
	Datetime  time.Time
}

type VIF struct {
	AllowedOperations      []string          `xmlrpc:"Allowed_operations"`
	CurrentOperations      map[string]string `xmlrpc:"Current_operations"`
	CurrentlyAttached      bool              `xmlrpc:"Currently_attached"`
	Device                 string
	MAC                    string
	MACAutogenerated       bool `xmlrpc:"MAC_autogenerated"`
	Metrics                string
	MTU                    int
	Network                string
	OtherConfig            map[string]string `xmlrpc:"Other_config"`
	QosAlgorithmParams     map[string]string `xmlrpc:"Qos_algorithm_params"`
	QosAlgorithmType       string            `xmlrpc:"Qos_algorithm_type"`
	QosSupportedAlgorithms []string          `xmlrpc:"Qos_supported_algorithms"`
	RuntimeProperties      map[string]string `xmlrpc:"Runtime_properties"`
	StatusCode             int               `xmlrpc:"Status_code"`
	StatusDetail           string            `xmlrpc:"Status_detail"`
	UUID                   string            `json:"uuid"`
	VM                     string
}

type PIF struct {
	BondMasterOf         []string `json:"Bond_master_of"`
	BondSlaveOf          string   `json:"Bond_slave_of"`
	CurrentlyAttached    bool     `json:"Currently_attached"`
	Device               string
	DisallowUnplug       bool `json:"Disallow_unplug"`
	DNS                  string
	Gateway              string
	Host                 string
	IP                   string
	IPConfigurationMode  string `json:"Ip_configuration_mode"`
	MAC                  string
	Management           bool
	Metrics              string
	MTU                  int
	Netmask              string
	Network              string
	OtherConfig          map[string]string `json:"Other_config"`
	Physical             bool
	TunnelAccessPIFOf    string `json:"Tunnel_access_PIF_of"`
	TunnelTransportPIFOf string `json:"Tunnel_transport_PIF_of"`
	UUID                 string `json:"uuid"`
	VLAN                 int
	VLANMasterOf         string   `json:"VLAN_master_of"`
	VLANSlaveOf          []string `json:"VLAN_slave_of"`
}

type Host struct {
	Address                        string
	AllowedOperations              []string          `json:"Allowed_operations"`
	CurrentOperations              map[string]string `json:"Current_operations"`
	APIVersionMajor                int               `json:"API_version_major"`
	APIVersionMinor                int               `json:"API_version_minor"`
	APIVersionVendor               string            `json:"API_version_vendor"`
	APIVersionVendorImplementation map[string]string `json:"API_version_vendor_implementation"`
	BiosStrings                    map[string]string `json:"Bios_strings"`
	Blobs                          map[string][]byte
	Capabilities                   []string
	ChipsetInfo                    map[string]string `json:"Chipset_info"`
	CPUConfiguration               map[string]string `json:"Cpu_configuration"`
	CPUInfo                        map[string]string `json:"Cpu_info"`
	CrashDumpSR                    string            `json:"Crash_dump_sr"`
	Crashdumps                     []string
	Edition                        string
	Enabled                        bool
	ExternalAuthConfiguration      map[string]string `json:"External_auth_configuration"`
	ExternalAuthType               string            `json:"External_auth_type"`
	ExternalAuthServiceName        string            `json:"External_auth_service_name"`
	HaNetworkPeers                 []string          `json:"Ha_network_peers"`
	HaStatefiles                   []string          `json:"Ha_statefiles"`
	HostCPUs                       []string          `json:"Host_CPUs"`
	Hostname                       string
	LicenseParams                  map[string]string `json:"License_params"`
	LicenseServer                  map[string]string `json:"License_server"`
	LocalCacheSR                   string            `json:"Local_cache_sr"`
	Logging                        map[string]string
	MemoryOverhead                 int `json:"Memory_overhead"`
	Metrics                        string
	NameDescription                string            `json:"Name_description"`
	NameLabel                      string            `json:"Name_label"`
	OtherConfig                    map[string]string `json:"Other_config"`
	Patches                        []string
	PBDs                           []string
	PCIs                           []string
	PGPUs                          []string
	PIFs                           []string
	PowerOnConfig                  map[string]string `json:"Power_on_config"`
	ResidentVMs                    []string          `json:"Resident_VMs"`
	SchedPolicy                    string            `json:"Sched_policy"`
	SoftwareVersion                map[string]string `json:"Software_version"`
	SupportedBootloaders           []string          `json:"Supported_bootloaders"`
	SuspendImageSR                 string            `json:"Suspend_image_sr"`
	UUID                           string            `json:"uuid"`
	Tags                           []string
}
