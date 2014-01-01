function CommonUtil() {
}

CommonUtil.prototype.getFuncOrString(text) {
	if (eval("typeof(" + text + ")") == "function") {
		return eval(text);
	}
	return text;
}
