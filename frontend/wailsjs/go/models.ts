export namespace config {
	
	export class AppConfig {
	    showHidden: boolean;
	    alwaysTop: boolean;
	    devMode: boolean;
	    autoConnect: boolean;
	    adbPath: string;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.showHidden = source["showHidden"];
	        this.alwaysTop = source["alwaysTop"];
	        this.devMode = source["devMode"];
	        this.autoConnect = source["autoConnect"];
	        this.adbPath = source["adbPath"];
	    }
	}

}

export namespace main {
	
	export class Device {
	    id: string;
	    model: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Device(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.model = source["model"];
	        this.status = source["status"];
	    }
	}
	export class DeviceProp {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new DeviceProp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class DeviceInfoResponse {
	    model: string;
	    manufacturer: string;
	    androidVer: string;
	    apiLevel: string;
	    cpu: string;
	    serial: string;
	    battery: string;
	    allProps: DeviceProp[];
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfoResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.model = source["model"];
	        this.manufacturer = source["manufacturer"];
	        this.androidVer = source["androidVer"];
	        this.apiLevel = source["apiLevel"];
	        this.cpu = source["cpu"];
	        this.serial = source["serial"];
	        this.battery = source["battery"];
	        this.allProps = this.convertValues(source["allProps"], DeviceProp);
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
	
	export class FileEntry {
	    name: string;
	    isDir: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isDir = source["isDir"];
	    }
	}
	export class ProcessInfo {
	    pid: string;
	    user: string;
	    cpu: string;
	    mem: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pid = source["pid"];
	        this.user = source["user"];
	        this.cpu = source["cpu"];
	        this.mem = source["mem"];
	        this.name = source["name"];
	    }
	}

}

