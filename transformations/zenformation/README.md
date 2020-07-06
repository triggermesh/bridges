# zenformation
Zendesk -> AWS Comprehend -> Zendesk

The following ENV var's must be present:

LANGUAGE

AWS_ACCESS_KEY_ID

AWS_SECRET_ACCESS_KEY

AWS_REGION


Expects A cloud event payload of :

type In struct {
	// ID is the Zendesk Ticket ID that will be pased to aws for comprenision
	ID          string  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}


Sends a Cloud Event with a playload of :

type Out struct {
	// ID is the Zendesk Ticket ID that will be pased to the ZD target for updates.
	ID  int64  `json:"id"`
	Tag string `json:"tag"`
}



Ex:

  curl -v https://zendesk-comprehiend.midimansland.dev.munu.io  \
    -H "Content-Type: application/json" \
    -H "Ce-Specversion: 1.0" \
    -H "Ce-Type: com.zendesk.ticket.create" \
    -H "Ce-Source: some.origin/intance" \
    -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
    -d '{"id": "86", "description" : "I love you you"}'  // or try ->   -d '{"id": 40, "description" : "I hate  you"}'


    
