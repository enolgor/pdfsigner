export namespace certs {
	
	export class StoredCertificateID {
	    Issuer: string;
	    Subject: string;
	
	    static createFrom(source: any = {}) {
	        return new StoredCertificateID(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Issuer = source["Issuer"];
	        this.Subject = source["Subject"];
	    }
	}

}

