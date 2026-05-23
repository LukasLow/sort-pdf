export namespace src {
	
	export class PdfInfo {
	    fileName: string;
	    sizeBytes: number;
	    pageCount: number;
	    dpi: number;
	
	    static createFrom(source: any = {}) {
	        return new PdfInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.sizeBytes = source["sizeBytes"];
	        this.pageCount = source["pageCount"];
	        this.dpi = source["dpi"];
	    }
	}

}

