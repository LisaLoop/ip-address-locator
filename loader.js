
const fs = require("fs");

fs.readFile("input.csv","utf8", function(err, contents){
	let lines = contents.split("\r\n");
	for(let i=0; i<lines.length;i++){
		var line = lines[i];
		line = line.replace(/'/gi,"");
		line = line.replace(/"/gi,"'");
                let sql = "INSERT INTO ipaddress(range_start,range_end,country_code,country,state,city) VALUES("+line+");";
		console.log(sql);
		//if(i===10){
		//	break;
		//}
	}
});
