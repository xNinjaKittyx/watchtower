package container

const (
	watchtowerLabel = "com.centurylinklabs.watchtower"
	signalLabel     = "com.centurylinklabs.watchtower.stop-signal"
	enableLabel     = "com.centurylinklabs.watchtower.enable"
	zodiacLabel     = "com.centurylinklabs.zodiac.original-image"
	preUpdateLabel  = "com.centurylinklabs.watchtower.lifecycle.pre-update"
	postUpdateLabel = "com.centurylinklabs.watchtower.lifecycle.post-update"
)

// GetLifecyclePreUpdateCommand returns the pre-update command set in the container metadata or an empty string
func (c Container) GetLifecyclePreUpdateCommand() string {
	return c.getLabelValueOrEmpty(preUpdateLabel)
}

// GetLifecyclePostUpdateCommand returns the post-update command set in the container metadata or an empty string
func (c Container) GetLifecyclePostUpdateCommand() string {
	return c.getLabelValueOrEmpty(postUpdateLabel)
}

// ContainsWatchtowerLabel takes a map of labels and values and tells
// the consumer whether it contains a valid watchtower instance label
func ContainsWatchtowerLabel(labels map[string]string) bool {
	val, ok := labels[watchtowerLabel]
	return ok && val == "true"
}

func (c Container) getLabelValueOrEmpty(label string) string {
	if val, ok := c.containerInfo.Config.Labels[label]; ok {
		return val
	}
	return ""
}

func (c Container) getLabelValue(label string) (string, bool) {
	val, ok := c.containerInfo.Config.Labels[label]
	return val, ok
}