db.createCollection('campaigns');
db.createCollection('votes');
db.votes.createIndex({campaignID: 1, participantID: 1}, {unique: true, name: "campaignIDParticipantIDUniqueIndex"});