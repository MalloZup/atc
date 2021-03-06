// This file was generated by counterfeiter
package workerfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeClient struct {
	FindOrCreateBuildContainerStub        func(lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, int, atc.PlanID, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes) (worker.Container, error)
	findOrCreateBuildContainerMutex       sync.RWMutex
	findOrCreateBuildContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 int
		arg5 atc.PlanID
		arg6 dbng.ContainerMetadata
		arg7 worker.ContainerSpec
		arg8 atc.VersionedResourceTypes
	}
	findOrCreateBuildContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	findOrCreateBuildContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 error
	}
	CreateResourceGetContainerStub        func(logger lager.Logger, resourceUser dbng.ResourceUser, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, metadata dbng.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error)
	createResourceGetContainerMutex       sync.RWMutex
	createResourceGetContainerArgsForCall []struct {
		logger        lager.Logger
		resourceUser  dbng.ResourceUser
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		metadata      dbng.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}
	createResourceGetContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	createResourceGetContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 error
	}
	FindOrCreateResourceCheckContainerStub        func(logger lager.Logger, resourceUser dbng.ResourceUser, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, metadata dbng.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, resourceType string, source atc.Source) (worker.Container, error)
	findOrCreateResourceCheckContainerMutex       sync.RWMutex
	findOrCreateResourceCheckContainerArgsForCall []struct {
		logger        lager.Logger
		resourceUser  dbng.ResourceUser
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		metadata      dbng.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
		resourceType  string
		source        atc.Source
	}
	findOrCreateResourceCheckContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	findOrCreateResourceCheckContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 error
	}
	CreateVolumeForResourceCacheStub        func(logger lager.Logger, vs worker.VolumeSpec, resourceCache *dbng.UsedResourceCache) (worker.Volume, error)
	createVolumeForResourceCacheMutex       sync.RWMutex
	createVolumeForResourceCacheArgsForCall []struct {
		logger        lager.Logger
		vs            worker.VolumeSpec
		resourceCache *dbng.UsedResourceCache
	}
	createVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 error
	}
	createVolumeForResourceCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 error
	}
	FindInitializedVolumeForResourceCacheStub        func(logger lager.Logger, resourceCache *dbng.UsedResourceCache) (worker.Volume, bool, error)
	findInitializedVolumeForResourceCacheMutex       sync.RWMutex
	findInitializedVolumeForResourceCacheArgsForCall []struct {
		logger        lager.Logger
		resourceCache *dbng.UsedResourceCache
	}
	findInitializedVolumeForResourceCacheReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findInitializedVolumeForResourceCacheReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindContainerByHandleStub        func(lager.Logger, int, string) (worker.Container, bool, error)
	findContainerByHandleMutex       sync.RWMutex
	findContainerByHandleArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}
	findContainerByHandleReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	findContainerByHandleReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	FindResourceTypeByPathStub        func(path string) (atc.WorkerResourceType, bool)
	findResourceTypeByPathMutex       sync.RWMutex
	findResourceTypeByPathArgsForCall []struct {
		path string
	}
	findResourceTypeByPathReturns struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	findResourceTypeByPathReturnsOnCall map[int]struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	lookupVolumeReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	SatisfyingStub        func(worker.WorkerSpec, atc.VersionedResourceTypes) (worker.Worker, error)
	satisfyingMutex       sync.RWMutex
	satisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.VersionedResourceTypes
	}
	satisfyingReturns struct {
		result1 worker.Worker
		result2 error
	}
	satisfyingReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	AllSatisfyingStub        func(worker.WorkerSpec, atc.VersionedResourceTypes) ([]worker.Worker, error)
	allSatisfyingMutex       sync.RWMutex
	allSatisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.VersionedResourceTypes
	}
	allSatisfyingReturns struct {
		result1 []worker.Worker
		result2 error
	}
	allSatisfyingReturnsOnCall map[int]struct {
		result1 []worker.Worker
		result2 error
	}
	RunningWorkersStub        func() ([]worker.Worker, error)
	runningWorkersMutex       sync.RWMutex
	runningWorkersArgsForCall []struct{}
	runningWorkersReturns     struct {
		result1 []worker.Worker
		result2 error
	}
	runningWorkersReturnsOnCall map[int]struct {
		result1 []worker.Worker
		result2 error
	}
	GetWorkerStub        func(workerName string) (worker.Worker, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		workerName string
	}
	getWorkerReturns struct {
		result1 worker.Worker
		result2 error
	}
	getWorkerReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) FindOrCreateBuildContainer(arg1 lager.Logger, arg2 <-chan os.Signal, arg3 worker.ImageFetchingDelegate, arg4 int, arg5 atc.PlanID, arg6 dbng.ContainerMetadata, arg7 worker.ContainerSpec, arg8 atc.VersionedResourceTypes) (worker.Container, error) {
	fake.findOrCreateBuildContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateBuildContainerReturnsOnCall[len(fake.findOrCreateBuildContainerArgsForCall)]
	fake.findOrCreateBuildContainerArgsForCall = append(fake.findOrCreateBuildContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 int
		arg5 atc.PlanID
		arg6 dbng.ContainerMetadata
		arg7 worker.ContainerSpec
		arg8 atc.VersionedResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("FindOrCreateBuildContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.findOrCreateBuildContainerMutex.Unlock()
	if fake.FindOrCreateBuildContainerStub != nil {
		return fake.FindOrCreateBuildContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateBuildContainerReturns.result1, fake.findOrCreateBuildContainerReturns.result2
}

func (fake *FakeClient) FindOrCreateBuildContainerCallCount() int {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return len(fake.findOrCreateBuildContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateBuildContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, int, atc.PlanID, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes) {
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	return fake.findOrCreateBuildContainerArgsForCall[i].arg1, fake.findOrCreateBuildContainerArgsForCall[i].arg2, fake.findOrCreateBuildContainerArgsForCall[i].arg3, fake.findOrCreateBuildContainerArgsForCall[i].arg4, fake.findOrCreateBuildContainerArgsForCall[i].arg5, fake.findOrCreateBuildContainerArgsForCall[i].arg6, fake.findOrCreateBuildContainerArgsForCall[i].arg7, fake.findOrCreateBuildContainerArgsForCall[i].arg8
}

func (fake *FakeClient) FindOrCreateBuildContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateBuildContainerStub = nil
	fake.findOrCreateBuildContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateBuildContainerReturnsOnCall(i int, result1 worker.Container, result2 error) {
	fake.FindOrCreateBuildContainerStub = nil
	if fake.findOrCreateBuildContainerReturnsOnCall == nil {
		fake.findOrCreateBuildContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 error
		})
	}
	fake.findOrCreateBuildContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateResourceGetContainer(logger lager.Logger, resourceUser dbng.ResourceUser, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, metadata dbng.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error) {
	fake.createResourceGetContainerMutex.Lock()
	ret, specificReturn := fake.createResourceGetContainerReturnsOnCall[len(fake.createResourceGetContainerArgsForCall)]
	fake.createResourceGetContainerArgsForCall = append(fake.createResourceGetContainerArgsForCall, struct {
		logger        lager.Logger
		resourceUser  dbng.ResourceUser
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		metadata      dbng.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}{logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, version, source, params})
	fake.recordInvocation("CreateResourceGetContainer", []interface{}{logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, version, source, params})
	fake.createResourceGetContainerMutex.Unlock()
	if fake.CreateResourceGetContainerStub != nil {
		return fake.CreateResourceGetContainerStub(logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, version, source, params)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createResourceGetContainerReturns.result1, fake.createResourceGetContainerReturns.result2
}

func (fake *FakeClient) CreateResourceGetContainerCallCount() int {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return len(fake.createResourceGetContainerArgsForCall)
}

func (fake *FakeClient) CreateResourceGetContainerArgsForCall(i int) (lager.Logger, dbng.ResourceUser, <-chan os.Signal, worker.ImageFetchingDelegate, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes, string, atc.Version, atc.Source, atc.Params) {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return fake.createResourceGetContainerArgsForCall[i].logger, fake.createResourceGetContainerArgsForCall[i].resourceUser, fake.createResourceGetContainerArgsForCall[i].cancel, fake.createResourceGetContainerArgsForCall[i].delegate, fake.createResourceGetContainerArgsForCall[i].metadata, fake.createResourceGetContainerArgsForCall[i].spec, fake.createResourceGetContainerArgsForCall[i].resourceTypes, fake.createResourceGetContainerArgsForCall[i].resourceType, fake.createResourceGetContainerArgsForCall[i].version, fake.createResourceGetContainerArgsForCall[i].source, fake.createResourceGetContainerArgsForCall[i].params
}

func (fake *FakeClient) CreateResourceGetContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateResourceGetContainerStub = nil
	fake.createResourceGetContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateResourceGetContainerReturnsOnCall(i int, result1 worker.Container, result2 error) {
	fake.CreateResourceGetContainerStub = nil
	if fake.createResourceGetContainerReturnsOnCall == nil {
		fake.createResourceGetContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 error
		})
	}
	fake.createResourceGetContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateResourceCheckContainer(logger lager.Logger, resourceUser dbng.ResourceUser, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, metadata dbng.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes, resourceType string, source atc.Source) (worker.Container, error) {
	fake.findOrCreateResourceCheckContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceCheckContainerReturnsOnCall[len(fake.findOrCreateResourceCheckContainerArgsForCall)]
	fake.findOrCreateResourceCheckContainerArgsForCall = append(fake.findOrCreateResourceCheckContainerArgsForCall, struct {
		logger        lager.Logger
		resourceUser  dbng.ResourceUser
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		metadata      dbng.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
		resourceType  string
		source        atc.Source
	}{logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, source})
	fake.recordInvocation("FindOrCreateResourceCheckContainer", []interface{}{logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, source})
	fake.findOrCreateResourceCheckContainerMutex.Unlock()
	if fake.FindOrCreateResourceCheckContainerStub != nil {
		return fake.FindOrCreateResourceCheckContainerStub(logger, resourceUser, cancel, delegate, metadata, spec, resourceTypes, resourceType, source)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateResourceCheckContainerReturns.result1, fake.findOrCreateResourceCheckContainerReturns.result2
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerCallCount() int {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return len(fake.findOrCreateResourceCheckContainerArgsForCall)
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerArgsForCall(i int) (lager.Logger, dbng.ResourceUser, <-chan os.Signal, worker.ImageFetchingDelegate, dbng.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes, string, atc.Source) {
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	return fake.findOrCreateResourceCheckContainerArgsForCall[i].logger, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceUser, fake.findOrCreateResourceCheckContainerArgsForCall[i].cancel, fake.findOrCreateResourceCheckContainerArgsForCall[i].delegate, fake.findOrCreateResourceCheckContainerArgsForCall[i].metadata, fake.findOrCreateResourceCheckContainerArgsForCall[i].spec, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceTypes, fake.findOrCreateResourceCheckContainerArgsForCall[i].resourceType, fake.findOrCreateResourceCheckContainerArgsForCall[i].source
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateResourceCheckContainerStub = nil
	fake.findOrCreateResourceCheckContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindOrCreateResourceCheckContainerReturnsOnCall(i int, result1 worker.Container, result2 error) {
	fake.FindOrCreateResourceCheckContainerStub = nil
	if fake.findOrCreateResourceCheckContainerReturnsOnCall == nil {
		fake.findOrCreateResourceCheckContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 error
		})
	}
	fake.findOrCreateResourceCheckContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateVolumeForResourceCache(logger lager.Logger, vs worker.VolumeSpec, resourceCache *dbng.UsedResourceCache) (worker.Volume, error) {
	fake.createVolumeForResourceCacheMutex.Lock()
	ret, specificReturn := fake.createVolumeForResourceCacheReturnsOnCall[len(fake.createVolumeForResourceCacheArgsForCall)]
	fake.createVolumeForResourceCacheArgsForCall = append(fake.createVolumeForResourceCacheArgsForCall, struct {
		logger        lager.Logger
		vs            worker.VolumeSpec
		resourceCache *dbng.UsedResourceCache
	}{logger, vs, resourceCache})
	fake.recordInvocation("CreateVolumeForResourceCache", []interface{}{logger, vs, resourceCache})
	fake.createVolumeForResourceCacheMutex.Unlock()
	if fake.CreateVolumeForResourceCacheStub != nil {
		return fake.CreateVolumeForResourceCacheStub(logger, vs, resourceCache)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createVolumeForResourceCacheReturns.result1, fake.createVolumeForResourceCacheReturns.result2
}

func (fake *FakeClient) CreateVolumeForResourceCacheCallCount() int {
	fake.createVolumeForResourceCacheMutex.RLock()
	defer fake.createVolumeForResourceCacheMutex.RUnlock()
	return len(fake.createVolumeForResourceCacheArgsForCall)
}

func (fake *FakeClient) CreateVolumeForResourceCacheArgsForCall(i int) (lager.Logger, worker.VolumeSpec, *dbng.UsedResourceCache) {
	fake.createVolumeForResourceCacheMutex.RLock()
	defer fake.createVolumeForResourceCacheMutex.RUnlock()
	return fake.createVolumeForResourceCacheArgsForCall[i].logger, fake.createVolumeForResourceCacheArgsForCall[i].vs, fake.createVolumeForResourceCacheArgsForCall[i].resourceCache
}

func (fake *FakeClient) CreateVolumeForResourceCacheReturns(result1 worker.Volume, result2 error) {
	fake.CreateVolumeForResourceCacheStub = nil
	fake.createVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateVolumeForResourceCacheReturnsOnCall(i int, result1 worker.Volume, result2 error) {
	fake.CreateVolumeForResourceCacheStub = nil
	if fake.createVolumeForResourceCacheReturnsOnCall == nil {
		fake.createVolumeForResourceCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 error
		})
	}
	fake.createVolumeForResourceCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindInitializedVolumeForResourceCache(logger lager.Logger, resourceCache *dbng.UsedResourceCache) (worker.Volume, bool, error) {
	fake.findInitializedVolumeForResourceCacheMutex.Lock()
	ret, specificReturn := fake.findInitializedVolumeForResourceCacheReturnsOnCall[len(fake.findInitializedVolumeForResourceCacheArgsForCall)]
	fake.findInitializedVolumeForResourceCacheArgsForCall = append(fake.findInitializedVolumeForResourceCacheArgsForCall, struct {
		logger        lager.Logger
		resourceCache *dbng.UsedResourceCache
	}{logger, resourceCache})
	fake.recordInvocation("FindInitializedVolumeForResourceCache", []interface{}{logger, resourceCache})
	fake.findInitializedVolumeForResourceCacheMutex.Unlock()
	if fake.FindInitializedVolumeForResourceCacheStub != nil {
		return fake.FindInitializedVolumeForResourceCacheStub(logger, resourceCache)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findInitializedVolumeForResourceCacheReturns.result1, fake.findInitializedVolumeForResourceCacheReturns.result2, fake.findInitializedVolumeForResourceCacheReturns.result3
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheCallCount() int {
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	return len(fake.findInitializedVolumeForResourceCacheArgsForCall)
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheArgsForCall(i int) (lager.Logger, *dbng.UsedResourceCache) {
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	return fake.findInitializedVolumeForResourceCacheArgsForCall[i].logger, fake.findInitializedVolumeForResourceCacheArgsForCall[i].resourceCache
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindInitializedVolumeForResourceCacheStub = nil
	fake.findInitializedVolumeForResourceCacheReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindInitializedVolumeForResourceCacheReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.FindInitializedVolumeForResourceCacheStub = nil
	if fake.findInitializedVolumeForResourceCacheReturnsOnCall == nil {
		fake.findInitializedVolumeForResourceCacheReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findInitializedVolumeForResourceCacheReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindContainerByHandle(arg1 lager.Logger, arg2 int, arg3 string) (worker.Container, bool, error) {
	fake.findContainerByHandleMutex.Lock()
	ret, specificReturn := fake.findContainerByHandleReturnsOnCall[len(fake.findContainerByHandleArgsForCall)]
	fake.findContainerByHandleArgsForCall = append(fake.findContainerByHandleArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindContainerByHandle", []interface{}{arg1, arg2, arg3})
	fake.findContainerByHandleMutex.Unlock()
	if fake.FindContainerByHandleStub != nil {
		return fake.FindContainerByHandleStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findContainerByHandleReturns.result1, fake.findContainerByHandleReturns.result2, fake.findContainerByHandleReturns.result3
}

func (fake *FakeClient) FindContainerByHandleCallCount() int {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return len(fake.findContainerByHandleArgsForCall)
}

func (fake *FakeClient) FindContainerByHandleArgsForCall(i int) (lager.Logger, int, string) {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return fake.findContainerByHandleArgsForCall[i].arg1, fake.findContainerByHandleArgsForCall[i].arg2, fake.findContainerByHandleArgsForCall[i].arg3
}

func (fake *FakeClient) FindContainerByHandleReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerByHandleStub = nil
	fake.findContainerByHandleReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindContainerByHandleReturnsOnCall(i int, result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerByHandleStub = nil
	if fake.findContainerByHandleReturnsOnCall == nil {
		fake.findContainerByHandleReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 bool
			result3 error
		})
	}
	fake.findContainerByHandleReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) FindResourceTypeByPath(path string) (atc.WorkerResourceType, bool) {
	fake.findResourceTypeByPathMutex.Lock()
	ret, specificReturn := fake.findResourceTypeByPathReturnsOnCall[len(fake.findResourceTypeByPathArgsForCall)]
	fake.findResourceTypeByPathArgsForCall = append(fake.findResourceTypeByPathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("FindResourceTypeByPath", []interface{}{path})
	fake.findResourceTypeByPathMutex.Unlock()
	if fake.FindResourceTypeByPathStub != nil {
		return fake.FindResourceTypeByPathStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findResourceTypeByPathReturns.result1, fake.findResourceTypeByPathReturns.result2
}

func (fake *FakeClient) FindResourceTypeByPathCallCount() int {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return len(fake.findResourceTypeByPathArgsForCall)
}

func (fake *FakeClient) FindResourceTypeByPathArgsForCall(i int) string {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return fake.findResourceTypeByPathArgsForCall[i].path
}

func (fake *FakeClient) FindResourceTypeByPathReturns(result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	fake.findResourceTypeByPathReturns = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeClient) FindResourceTypeByPathReturnsOnCall(i int, result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	if fake.findResourceTypeByPathReturnsOnCall == nil {
		fake.findResourceTypeByPathReturnsOnCall = make(map[int]struct {
			result1 atc.WorkerResourceType
			result2 bool
		})
	}
	fake.findResourceTypeByPathReturnsOnCall[i] = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	ret, specificReturn := fake.lookupVolumeReturnsOnCall[len(fake.lookupVolumeArgsForCall)]
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
}

func (fake *FakeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) LookupVolumeReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	if fake.lookupVolumeReturnsOnCall == nil {
		fake.lookupVolumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.lookupVolumeReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Satisfying(arg1 worker.WorkerSpec, arg2 atc.VersionedResourceTypes) (worker.Worker, error) {
	fake.satisfyingMutex.Lock()
	ret, specificReturn := fake.satisfyingReturnsOnCall[len(fake.satisfyingArgsForCall)]
	fake.satisfyingArgsForCall = append(fake.satisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.VersionedResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("Satisfying", []interface{}{arg1, arg2})
	fake.satisfyingMutex.Unlock()
	if fake.SatisfyingStub != nil {
		return fake.SatisfyingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.satisfyingReturns.result1, fake.satisfyingReturns.result2
}

func (fake *FakeClient) SatisfyingCallCount() int {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return len(fake.satisfyingArgsForCall)
}

func (fake *FakeClient) SatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.VersionedResourceTypes) {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return fake.satisfyingArgsForCall[i].arg1, fake.satisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) SatisfyingReturns(result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	fake.satisfyingReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) SatisfyingReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	if fake.satisfyingReturnsOnCall == nil {
		fake.satisfyingReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.satisfyingReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AllSatisfying(arg1 worker.WorkerSpec, arg2 atc.VersionedResourceTypes) ([]worker.Worker, error) {
	fake.allSatisfyingMutex.Lock()
	ret, specificReturn := fake.allSatisfyingReturnsOnCall[len(fake.allSatisfyingArgsForCall)]
	fake.allSatisfyingArgsForCall = append(fake.allSatisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.VersionedResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("AllSatisfying", []interface{}{arg1, arg2})
	fake.allSatisfyingMutex.Unlock()
	if fake.AllSatisfyingStub != nil {
		return fake.AllSatisfyingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allSatisfyingReturns.result1, fake.allSatisfyingReturns.result2
}

func (fake *FakeClient) AllSatisfyingCallCount() int {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return len(fake.allSatisfyingArgsForCall)
}

func (fake *FakeClient) AllSatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.VersionedResourceTypes) {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return fake.allSatisfyingArgsForCall[i].arg1, fake.allSatisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) AllSatisfyingReturns(result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	fake.allSatisfyingReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AllSatisfyingReturnsOnCall(i int, result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	if fake.allSatisfyingReturnsOnCall == nil {
		fake.allSatisfyingReturnsOnCall = make(map[int]struct {
			result1 []worker.Worker
			result2 error
		})
	}
	fake.allSatisfyingReturnsOnCall[i] = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunningWorkers() ([]worker.Worker, error) {
	fake.runningWorkersMutex.Lock()
	ret, specificReturn := fake.runningWorkersReturnsOnCall[len(fake.runningWorkersArgsForCall)]
	fake.runningWorkersArgsForCall = append(fake.runningWorkersArgsForCall, struct{}{})
	fake.recordInvocation("RunningWorkers", []interface{}{})
	fake.runningWorkersMutex.Unlock()
	if fake.RunningWorkersStub != nil {
		return fake.RunningWorkersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runningWorkersReturns.result1, fake.runningWorkersReturns.result2
}

func (fake *FakeClient) RunningWorkersCallCount() int {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return len(fake.runningWorkersArgsForCall)
}

func (fake *FakeClient) RunningWorkersReturns(result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	fake.runningWorkersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RunningWorkersReturnsOnCall(i int, result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	if fake.runningWorkersReturnsOnCall == nil {
		fake.runningWorkersReturnsOnCall = make(map[int]struct {
			result1 []worker.Worker
			result2 error
		})
	}
	fake.runningWorkersReturnsOnCall[i] = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetWorker(workerName string) (worker.Worker, error) {
	fake.getWorkerMutex.Lock()
	ret, specificReturn := fake.getWorkerReturnsOnCall[len(fake.getWorkerArgsForCall)]
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("GetWorker", []interface{}{workerName})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(workerName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2
}

func (fake *FakeClient) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeClient) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].workerName
}

func (fake *FakeClient) GetWorkerReturns(result1 worker.Worker, result2 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetWorkerReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.GetWorkerStub = nil
	if fake.getWorkerReturnsOnCall == nil {
		fake.getWorkerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.getWorkerReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateBuildContainerMutex.RLock()
	defer fake.findOrCreateBuildContainerMutex.RUnlock()
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	fake.findOrCreateResourceCheckContainerMutex.RLock()
	defer fake.findOrCreateResourceCheckContainerMutex.RUnlock()
	fake.createVolumeForResourceCacheMutex.RLock()
	defer fake.createVolumeForResourceCacheMutex.RUnlock()
	fake.findInitializedVolumeForResourceCacheMutex.RLock()
	defer fake.findInitializedVolumeForResourceCacheMutex.RUnlock()
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.Client = new(FakeClient)
