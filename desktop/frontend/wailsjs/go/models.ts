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

export namespace color {
	
	export class RGBA {
	    R: number;
	    G: number;
	    B: number;
	    A: number;
	
	    static createFrom(source: any = {}) {
	        return new RGBA(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.R = source["R"];
	        this.G = source["G"];
	        this.B = source["B"];
	        this.A = source["A"];
	    }
	}

}

export namespace config {
	
	export enum Rotation {
	    ROTATE_0 = "0",
	    ROTATE_90 = "90",
	    ROTATE_180 = "180",
	    ROTATE_270 = "270",
	}
	export enum Alignment {
	    LEFT = "left",
	    CENTER = "center",
	    RIGHT = "right",
	}
	export class TextLine {
	    Key: string;
	    Value: string;
	
	    static createFrom(source: any = {}) {
	        return new TextLine(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Key = source["Key"];
	        this.Value = source["Value"];
	    }
	}

}

export namespace stamps {
	
	export class StampConfig {
	    title: string;
	    dateFormat: string;
	    includeTitle: boolean;
	    includeSubject: boolean;
	    includeIssuer: boolean;
	    includeDate: boolean;
	    subjectKey: string;
	    issuerKey: string;
	    dateKey: string;
	    extraLines?: config.TextLine[];
	    dpi: number;
	    backgroundColor: color.RGBA;
	    widthPt: number;
	    heightPt: number;
	    posXPt: number;
	    posYPt: number;
	    posStrict: boolean;
	    rotate: config.Rotation;
	    borderSizePt: number;
	    borderColor: color.RGBA;
	    logo?: string;
	    logoOpacity: number;
	    logoGrayScale: boolean;
	    logoAlignment: config.Alignment;
	    emptyLineAfterTitle: boolean;
	    titleAlignment: config.Alignment;
	    lineAlignment: config.Alignment;
	    keyAlignment: config.Alignment;
	    valueAlignment: config.Alignment;
	    titleFont: string;
	    keyFont: string;
	    valueFont: string;
	    titleColor: color.RGBA;
	    keyColor: color.RGBA;
	    valueColor: color.RGBA;
	
	    static createFrom(source: any = {}) {
	        return new StampConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.dateFormat = source["dateFormat"];
	        this.includeTitle = source["includeTitle"];
	        this.includeSubject = source["includeSubject"];
	        this.includeIssuer = source["includeIssuer"];
	        this.includeDate = source["includeDate"];
	        this.subjectKey = source["subjectKey"];
	        this.issuerKey = source["issuerKey"];
	        this.dateKey = source["dateKey"];
	        this.extraLines = this.convertValues(source["extraLines"], config.TextLine);
	        this.dpi = source["dpi"];
	        this.backgroundColor = this.convertValues(source["backgroundColor"], color.RGBA);
	        this.widthPt = source["widthPt"];
	        this.heightPt = source["heightPt"];
	        this.posXPt = source["posXPt"];
	        this.posYPt = source["posYPt"];
	        this.posStrict = source["posStrict"];
	        this.rotate = source["rotate"];
	        this.borderSizePt = source["borderSizePt"];
	        this.borderColor = this.convertValues(source["borderColor"], color.RGBA);
	        this.logo = source["logo"];
	        this.logoOpacity = source["logoOpacity"];
	        this.logoGrayScale = source["logoGrayScale"];
	        this.logoAlignment = source["logoAlignment"];
	        this.emptyLineAfterTitle = source["emptyLineAfterTitle"];
	        this.titleAlignment = source["titleAlignment"];
	        this.lineAlignment = source["lineAlignment"];
	        this.keyAlignment = source["keyAlignment"];
	        this.valueAlignment = source["valueAlignment"];
	        this.titleFont = source["titleFont"];
	        this.keyFont = source["keyFont"];
	        this.valueFont = source["valueFont"];
	        this.titleColor = this.convertValues(source["titleColor"], color.RGBA);
	        this.keyColor = this.convertValues(source["keyColor"], color.RGBA);
	        this.valueColor = this.convertValues(source["valueColor"], color.RGBA);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

