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
	export class MoveConflict {
	    hasConflict: boolean;
	    sourcePath: string;
	    targetPath: string;
	    targetName: string;
	
	    static createFrom(source: any = {}) {
	        return new MoveConflict(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hasConflict = source["hasConflict"];
	        this.sourcePath = source["sourcePath"];
	        this.targetPath = source["targetPath"];
	        this.targetName = source["targetName"];
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

