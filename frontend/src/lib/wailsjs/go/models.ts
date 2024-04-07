export namespace kuddle {
	
	
	
	export class Session {
	    name: string;
	    access_key_id: string;
	    endpoint: string;
	    secret_access_key: string;
	    token: string;
	    use_ssl: boolean;
	    bucket: string;
	    domain: string;
	
	    static createFrom(source: any = {}) {
	        return new Session(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.access_key_id = source["access_key_id"];
	        this.endpoint = source["endpoint"];
	        this.secret_access_key = source["secret_access_key"];
	        this.token = source["token"];
	        this.use_ssl = source["use_ssl"];
	        this.bucket = source["bucket"];
	        this.domain = source["domain"];
	    }
	}

}

