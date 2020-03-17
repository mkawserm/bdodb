package bdodb

import "testing"

func TestDefaultLog_Debugf(t *testing.T) {
	DefaultLogger.Debugf("MSG")
}

func TestDefaultLog_Errorf(t *testing.T) {
	DefaultLogger.Errorf("MSG")
}

func TestDefaultLog_Infof(t *testing.T) {
	DefaultLogger.Infof("MSG")
}

func TestDefaultLog_Warningf(t *testing.T) {
	DefaultLogger.Warningf("MSG")
}
