package tests

//go:generate mockgen -package=mock_interfaces -destination=./mock_interfaces/fake_controller.go github.com/mohsin123321/cloud-project/controller ControllerInterface
//go:generate mockgen -package=mock_interfaces -destination=./mock_interfaces/fake_services.go github.com/mohsin123321/cloud-project/dataservice DataserviceInterface
//go:generate mockgen -package=mock_interfaces -destination=./mock_interfaces/fake_database.go github.com/mohsin123321/cloud-project/database DatabaseInterface
//go:generate mockgen -package=mock_interfaces -destination=./mock_interfaces/fake_utility.go github.com/mohsin123321/cloud-project/utility UtilityInterface
