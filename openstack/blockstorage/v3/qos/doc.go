/*
Package qos provides information and interaction with the QoS specifications
for the Openstack Blockstorage service.

Example to create a QoS specification

	createOpts := qos.CreateOpts{
		Name:     "test",
		Consumer: qos.ConsumerFront,
		Specs: map[string]string{
			"read_iops_sec": "20000",
		},
	}

	test, err := qos.Create(client, createOpts).Extract()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("QoS: %+v\n", test)

Example to delete a QoS specification

	qosID := "d6ae28ce-fcb5-4180-aa62-d260a27e09ae"

	deleteOpts := qos.DeleteOpts{
		Force: false,
	}

	err = qos.Delete(client, qosID, deleteOpts).ExtractErr()
	if err != nil {
		log.Fatal(err)
	}

Example to list QoS specifications

	listOpts := qos.ListOpts{}

	allPages, err := qos.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allQoS, err := qos.ExtractQoS(allPages)
	if err != nil {
		panic(err)
	}

	for _, qos := range allQoS {
		fmt.Printf("List: %+v\n", qos)
	}

Example to get a single QoS specification

	qosID := "de075d5e-8afc-4e23-9388-b84a5183d1c0"

	singleQos, err := qos.Get(client, test.ID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get: %+v\n", singleQos)

Example of updating QoSSpec

	qosID := "de075d5e-8afc-4e23-9388-b84a5183d1c0"

	updateOpts := qos.UpdateOpts{
		Consumer: qos.ConsumerBack,
		Specs: map[string]string{
			"read_iops_sec": "40000",
		},
	}

	specs, err := qos.Update(client, qosID, qosSpecs).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", specs)

*/
package qos
