db.createCollection('campaigns');
db.createCollection('votes');
db.votes.createIndex({campaign_id: 1, participant_id: 1}, {unique: true, name: "campaignIDParticipantIDUniqueIndex"});
db.participants.createIndex({passport_id: 1}, {unique: true, name: "participantIDUniqueIndex"});