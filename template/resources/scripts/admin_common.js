

var common = {
	view: {},
	system: {}
};

common.login = function(accessCode, account, passWord, on_result) {
	$.post("/login/", {
		accessCode: accessCode,
		account: account,
		password: passWord
	}, function(ret){
		console.log("common.login: ret: %s", JSON.stringify(ret));
		
		try{
			if (ret.ErrCode === 0) {
				common.system.account = ret.Account;
				common.system.accessCode = ret.AccessCode;
				
				on_result(1, "ok");
			} else {
				on_result(0, ret.cause);
			}
		}
		catch(e) {
			on_result(0, e.message);
		}
	}, "json");
};

common.logout = function(accessCode, on_result){
	$.post("/logout/", {
		accessCode: accessCode
	}, function(ret){
		console.log("common.logout: ret: %s", JSON.stringify(ret));
		
		try{
			if (ret.ErrCode === 0) {
				on_result(1, "ok");
			} else {
				on_result(0, ret.cause);
			}
		}
		catch(e) {
			on_result(0, e.message);
		}
	}, "json");
};




