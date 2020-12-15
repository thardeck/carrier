package kubernetes

type DeploymentID string

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Deployment
type Deployment interface {
	Deploy(Cluster) error
	Upgrade(Cluster) error
	SetDomain(d string)
	GetDomain() string
	Delete(Cluster) error
	Describe() string
	GetVersion() string
	NeededOptions() InstallationOptions
	Restore(Cluster, string) error
	Backup(Cluster, string) error
	ID() DeploymentID
}

// A list of deployment that should be installed together
type Installer struct {
	Deployments   []Deployment
	NeededOptions InstallationOptions
}

func NewInstaller(deployments ...Deployment) *Installer {
	return &Installer{
		Deployments: deployments,
	}
}

// GatherNeededOptions merges all options from all deployments.
// Also ignores Values set on shared options (to avoid Deployments setting
// values for other Deployments)
func (i *Installer) GatherNeededOptions() {
	for _, d := range i.Deployments {
		curatedOptions := InstallationOptions{}
		for _, opt := range d.NeededOptions() {
			newOpt := opt
			if opt.DeploymentID == "" {
				newOpt.Value = ""
			}
			curatedOptions = append(curatedOptions, newOpt)
		}

		i.NeededOptions = i.NeededOptions.Merge(curatedOptions)
	}
}

type OptionsReader interface {
	Read(InstallationOption) interface{}
}

// PopulateNeededOptions will try to give values to the needed options
// using the given OptionsReader. If none is given, the default is the
// InteractiveOptionsReader which will ask in the terminal.
// This method only populates what is possible and leaves the rest empty.
// TODO: Implement another method to validate that all options have been set.
func (i *Installer) PopulateNeededOptions(reader OptionsReader) {
	// if reader == nil {
	// 	reader = InteractiveOptionsReader // stdin
	// }
	// for aech optino ask it from the reader calling Read()
	// for _, d := range i.Deployments {
	// 	i.NeededOptions = i.NeededOptions.Merge(d.NeededOptions())
	// }
}

func (i *Installer) Install(cluster Cluster, options InstallationOptions) error {
	// fmt.Println(d.Describe())
	//	for _, := range i {

	// // Automatically set a deployment domain based on platform reported ExternalIPs
	// if d.GetDomain() == "" {
	// 	ips := cluster.GetPlatform().ExternalIPs()
	// 	if len(ips) == 0 {
	// 		return errors.New("Could not detect cluster ExternalIPs and no deployment domain was specified")
	// 	}
	// 	d.SetDomain(fmt.Sprintf("%s.nip.io", ips[0]))
	// }
	// return d.Deploy(cluster)
	return nil
}

func (i *Installer) Delete(cluster Cluster) error {
	//return d.Delete(cluster)
	return nil
}

func (i *Installer) Upgrade(cluster Cluster) error {
	//return d.Upgrade(cluster)
	return nil
}

func (i *Installer) Backup(cluster Cluster, output string) error {
	//return d.Backup(cluster, output)
	return nil
}

func (i *Installer) Restore(cluster Cluster, output string) error {
	//return d.Restore(cluster, output)
	return nil
}
