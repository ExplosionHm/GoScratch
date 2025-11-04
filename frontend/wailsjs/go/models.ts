export namespace project {
	
	export class NodeAppearance {
	    Color: number;
	
	    static createFrom(source: any = {}) {
	        return new NodeAppearance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Color = source["Color"];
	    }
	}
	export class Project {
	    name: string;
	    id: string;
	    owner: string;
	    collaborators: string[];
	    language: string;
	    libraries: Record<string, number>;
	    theme: number;
	    nodes: NodeAppearance[];
	    assets: string[];
	    Program: number[];
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.id = source["id"];
	        this.owner = source["owner"];
	        this.collaborators = source["collaborators"];
	        this.language = source["language"];
	        this.libraries = source["libraries"];
	        this.theme = source["theme"];
	        this.nodes = this.convertValues(source["nodes"], NodeAppearance);
	        this.assets = source["assets"];
	        this.Program = source["Program"];
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

