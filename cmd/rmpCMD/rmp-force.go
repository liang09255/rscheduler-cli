package rmpCMD

var force bool

func init() {
	RmpCMD.Flags().BoolVarP(&force, "force", "f", false, "force close")
}
