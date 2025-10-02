package repository

import (
	"failiverCheck/internal/app/ds"
	"fmt"
	"math"

	"github.com/google/uuid"
)

func GenerateFileName() (string, error) {
	fileName, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return fileName.String(), nil
}

func CreateNewFilePath(filePath string) (string, error) {
	fileName, err := GenerateFileName()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", filePath, fileName), nil
}

func calculateAvailable(sys_cacl *ds.SystemCalculation) (float32, error) {
	var systemAvailable = float32(1)
	if len(sys_cacl.ComponentsToSystemCalc) == 0 {
		return 0, fmt.Errorf("ComponentsToSystemCalc is empty")
	}
	for _, item := range sys_cacl.ComponentsToSystemCalc {
		component := item.Component
		available := component.Available
		if available == 0.0 {
			if component.MTBF == 0 || component.MTBF+component.MTTR == 0 {
				return 0, fmt.Errorf("wrong value for mtbf or mttr")
			}
			available = float32(component.MTBF) / float32(component.MTBF+component.MTTR)
		}
		componentAvailable := 1 - math.Pow(float64(1-available), float64(item.ReplicationCount))
		systemAvailable *= float32(componentAvailable)
	}
	return systemAvailable, nil
}
