export namespace v1 {
	
	export class ConfigMapKeySelector {
	    name?: string;
	    key: string;
	    optional?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ConfigMapKeySelector(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.key = source["key"];
	        this.optional = source["optional"];
	    }
	}
	export class SecretKeySelector {
	    name?: string;
	    key: string;
	    optional?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SecretKeySelector(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.key = source["key"];
	        this.optional = source["optional"];
	    }
	}
	export class ResourceFieldSelector {
	    containerName?: string;
	    resource: string;
	    // Go type: resource
	    divisor?: any;
	
	    static createFrom(source: any = {}) {
	        return new ResourceFieldSelector(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.containerName = source["containerName"];
	        this.resource = source["resource"];
	        this.divisor = this.convertValues(source["divisor"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class ObjectFieldSelector {
	    apiVersion?: string;
	    fieldPath: string;
	
	    static createFrom(source: any = {}) {
	        return new ObjectFieldSelector(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.apiVersion = source["apiVersion"];
	        this.fieldPath = source["fieldPath"];
	    }
	}
	export class EnvVarSource {
	    fieldRef?: ObjectFieldSelector;
	    resourceFieldRef?: ResourceFieldSelector;
	    configMapKeyRef?: ConfigMapKeySelector;
	    secretKeyRef?: SecretKeySelector;
	
	    static createFrom(source: any = {}) {
	        return new EnvVarSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fieldRef = this.convertValues(source["fieldRef"], ObjectFieldSelector);
	        this.resourceFieldRef = this.convertValues(source["resourceFieldRef"], ResourceFieldSelector);
	        this.configMapKeyRef = this.convertValues(source["configMapKeyRef"], ConfigMapKeySelector);
	        this.secretKeyRef = this.convertValues(source["secretKeyRef"], SecretKeySelector);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class EnvVar {
	    name: string;
	    value?: string;
	    valueFrom?: EnvVarSource;
	
	    static createFrom(source: any = {}) {
	        return new EnvVar(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.valueFrom = this.convertValues(source["valueFrom"], EnvVarSource);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

