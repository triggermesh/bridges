package zenformation

// In is the definition of the expected incoming data
type In struct {
	// ID is the Zendesk Ticket ID that will be pased to aws
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Out is the definition of how the data is expected to look going into Zendesk
type Out struct {
	// ID is the Zendesk Ticket ID that will be pased to the ZD target for updates.
	ID  int64  `json:"id"`
	Tag string `json:"tag"`
}

//  a Zendesk target expects Event's of type `com.zendesk.tag.create` Expect both a `id` and `tag` to be preset.
//   - **Example of type : `com.zendesk.tag.create`**
//     ```sh
//     curl -v https://zendesktarget-triggermesh-zendesk.jnlasersolutions.dev.munu.io  \
//     -H "Content-Type: application/json" \
//     -H "Ce-Specversion: 1.0" \
//     -H "Ce-Type: com.zendesk.tag.create" \
//     -H "Ce-Source: some.origin/intance" \
//     -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
//     -d '{"id":81 , "tag":"triggermesh"}'
//     ```
