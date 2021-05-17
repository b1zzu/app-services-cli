// autogenerated using
// jq '.. | .operationId?' openapi/srs-fleet-manager.json  | grep -v null

module.exports = {
  getContentById: async (c, req, res) => {
    res.status(200).json({
      "openapi": "3.0.2",
      "info": {
        "title": "Empty API",
        "version": "1.0.0",
        "description": "An example API design using OpenAPI."
      }
    });
  },
  listArtifactsInGroup: async (c, req, res) => {
    res.status(200).json({
      "artifacts": [
        {
          "groupId": "My-Group",
          "id": "Procurement-Invoice",
          "name": "Artifact Name",
          "description": "Description of the artifact",
          "labels": [
            "current",
            "internal"
          ],
          "type": "AVRO",
          "state": "ENABLED",
          "createdBy": "user1",
          "createdOn": "2019-03-22T12:51:19Z"
        }
      ],
      "count": 0
    });
  },
  createArtifact: async (c, req, res) => {
    res.status(200).json({
      "groupId": "My-Group",
      "id": "Procurement-Invoice",
      "name": "Artifact Name",
      "description": "Description of the artifact",
      "type": "AVRO",
      "version": 18,
      "createdBy": "user1",
      "createdOn": "2019-03-22T12:51:19Z",
      "modifiedBy": "user2",
      "modifiedOn": "2019-07-19T15:09:00Z",
      "globalId": 12984719247,
      "contentId": 82736,
      "labels": [
        "label-1",
        "label-2"
      ],
      "properties": {
        "custom-1": "foo",
        "custom-2": "bar"
      }
    });
  },
  deleteArtifact: async (c, req, res) => {
    res.status(204).json();
  },
  
  // Handling auth
  notFound: async (c, req, res) => res.status(404).json({ err: "not found" }),
  unauthorizedHandler: async (c, req, res) =>
    res.status(401).json({ err: "unauthorized" }),
};