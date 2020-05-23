package godocmodule

import (
	"testing"
	"os"
)

func TestGoDoc(t *testing.T) {
	
	if _, err := os.Stat("../out/docs/my-docs.html"); err != nil {
    		t.Errorf("File or path doesn't exist, error: %s",err);  
  	}
	
}
