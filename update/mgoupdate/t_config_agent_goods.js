conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_agent_goods').drop();

db.getCollection('t_agent_goods').insert({
    "_id" : ObjectId("58aa56dd6fec9cba7bbab63d"),
    "name" : "450张房卡",
    "goodsid" : 1,
    "description" : "价值450元人民币",
    "type" : "房卡",
    "num" : 450.0,
    "givenum" : 0.0,
    "price" : 500.0
});
