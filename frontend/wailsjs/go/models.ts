export namespace src {
	
	export class Analysis {
	    correspondent: string;
	
	    static createFrom(source: any = {}) {
	        return new Analysis(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.correspondent = source["correspondent"];
	    }
	}
	export class PdfInfo {
	    fileName: string;
	    sizeBytes: number;
	    pageCount: number;
	
	    static createFrom(source: any = {}) {
	        return new PdfInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.sizeBytes = source["sizeBytes"];
	        this.pageCount = source["pageCount"];
	    }
	}

}

