export namespace main {
	
	export class Settings {
	    theme: string;
	    keyBindings: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	        this.keyBindings = source["keyBindings"];
	    }
	}

}

