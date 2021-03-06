// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mApiSnapshotIn  = stats.Int64("api.gloo.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mApiSnapshotOut = stats.Int64("api.gloo.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")

	apisnapshotInView = &view.View{
		Name:        "api.gloo.solo.io_snap_emitter/snap_in",
		Measure:     mApiSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	apisnapshotOutView = &view.View{
		Name:        "api.gloo.solo.io/snap_emitter/snap_out",
		Measure:     mApiSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(apisnapshotInView, apisnapshotOutView)
}

type ApiEmitter interface {
	Register() error
	Artifact() ArtifactClient
	Endpoint() EndpointClient
	Proxy() ProxyClient
	Secret() SecretClient
	Upstream() UpstreamClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error)
}

func NewApiEmitter(artifactClient ArtifactClient, endpointClient EndpointClient, proxyClient ProxyClient, secretClient SecretClient, upstreamClient UpstreamClient) ApiEmitter {
	return NewApiEmitterWithEmit(artifactClient, endpointClient, proxyClient, secretClient, upstreamClient, make(chan struct{}))
}

func NewApiEmitterWithEmit(artifactClient ArtifactClient, endpointClient EndpointClient, proxyClient ProxyClient, secretClient SecretClient, upstreamClient UpstreamClient, emit <-chan struct{}) ApiEmitter {
	return &apiEmitter{
		artifact:  artifactClient,
		endpoint:  endpointClient,
		proxy:     proxyClient,
		secret:    secretClient,
		upstream:  upstreamClient,
		forceEmit: emit,
	}
}

type apiEmitter struct {
	forceEmit <-chan struct{}
	artifact  ArtifactClient
	endpoint  EndpointClient
	proxy     ProxyClient
	secret    SecretClient
	upstream  UpstreamClient
}

func (c *apiEmitter) Register() error {
	if err := c.artifact.Register(); err != nil {
		return err
	}
	if err := c.endpoint.Register(); err != nil {
		return err
	}
	if err := c.proxy.Register(); err != nil {
		return err
	}
	if err := c.secret.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	return nil
}

func (c *apiEmitter) Artifact() ArtifactClient {
	return c.artifact
}

func (c *apiEmitter) Endpoint() EndpointClient {
	return c.endpoint
}

func (c *apiEmitter) Proxy() ProxyClient {
	return c.proxy
}

func (c *apiEmitter) Secret() SecretClient {
	return c.secret
}

func (c *apiEmitter) Upstream() UpstreamClient {
	return c.upstream
}

func (c *apiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Artifact */
	type artifactListWithNamespace struct {
		list      ArtifactList
		namespace string
	}
	artifactChan := make(chan artifactListWithNamespace)
	/* Create channel for Endpoint */
	type endpointListWithNamespace struct {
		list      EndpointList
		namespace string
	}
	endpointChan := make(chan endpointListWithNamespace)
	/* Create channel for Proxy */
	type proxyListWithNamespace struct {
		list      ProxyList
		namespace string
	}
	proxyChan := make(chan proxyListWithNamespace)
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for Artifact */
		artifactNamespacesChan, artifactErrs, err := c.artifact.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Artifact watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, artifactErrs, namespace+"-artifacts")
		}(namespace)
		/* Setup watch for Endpoint */
		endpointNamespacesChan, endpointErrs, err := c.endpoint.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Endpoint watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, endpointErrs, namespace+"-endpoints")
		}(namespace)
		/* Setup watch for Proxy */
		proxyNamespacesChan, proxyErrs, err := c.proxy.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Proxy watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, proxyErrs, namespace+"-proxies")
		}(namespace)
		/* Setup watch for Secret */
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)
		/* Setup watch for Upstream */
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case artifactList := <-artifactNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case artifactChan <- artifactListWithNamespace{list: artifactList, namespace: namespace}:
					}
				case endpointList := <-endpointNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case endpointChan <- endpointListWithNamespace{list: endpointList, namespace: namespace}:
					}
				case proxyList := <-proxyNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case proxyChan <- proxyListWithNamespace{list: proxyList, namespace: namespace}:
					}
				case secretList := <-secretNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *ApiSnapshot)
	go func() {
		originalSnapshot := ApiSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mApiSnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      artifactNamespacedList := <- artifactChan
		      currentSnapshot.Artifacts.Clear(artifactNamespacedList.namespace)
		      artifactList := artifactNamespacedList.list
		   	currentSnapshot.Artifacts.Add(artifactList...)
		      endpointNamespacedList := <- endpointChan
		      currentSnapshot.Endpoints.Clear(endpointNamespacedList.namespace)
		      endpointList := endpointNamespacedList.list
		   	currentSnapshot.Endpoints.Add(endpointList...)
		      proxyNamespacedList := <- proxyChan
		      currentSnapshot.Proxies.Clear(proxyNamespacedList.namespace)
		      proxyList := proxyNamespacedList.list
		   	currentSnapshot.Proxies.Add(proxyList...)
		      secretNamespacedList := <- secretChan
		      currentSnapshot.Secrets.Clear(secretNamespacedList.namespace)
		      secretList := secretNamespacedList.list
		   	currentSnapshot.Secrets.Add(secretList...)
		      upstreamNamespacedList := <- upstreamChan
		      currentSnapshot.Upstreams.Clear(upstreamNamespacedList.namespace)
		      upstreamList := upstreamNamespacedList.list
		   	currentSnapshot.Upstreams.Add(upstreamList...)
		   		}
		*/

		for {
			record := func() { stats.Record(ctx, mApiSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case artifactNamespacedList := <-artifactChan:
				record()

				namespace := artifactNamespacedList.namespace
				artifactList := artifactNamespacedList.list

				currentSnapshot.Artifacts.Clear(namespace)
				currentSnapshot.Artifacts.Add(artifactList...)
			case endpointNamespacedList := <-endpointChan:
				record()

				namespace := endpointNamespacedList.namespace
				endpointList := endpointNamespacedList.list

				currentSnapshot.Endpoints.Clear(namespace)
				currentSnapshot.Endpoints.Add(endpointList...)
			case proxyNamespacedList := <-proxyChan:
				record()

				namespace := proxyNamespacedList.namespace
				proxyList := proxyNamespacedList.list

				currentSnapshot.Proxies.Clear(namespace)
				currentSnapshot.Proxies.Add(proxyList...)
			case secretNamespacedList := <-secretChan:
				record()

				namespace := secretNamespacedList.namespace
				secretList := secretNamespacedList.list

				currentSnapshot.Secrets.Clear(namespace)
				currentSnapshot.Secrets.Add(secretList...)
			case upstreamNamespacedList := <-upstreamChan:
				record()

				namespace := upstreamNamespacedList.namespace
				upstreamList := upstreamNamespacedList.list

				currentSnapshot.Upstreams.Clear(namespace)
				currentSnapshot.Upstreams.Add(upstreamList...)
			}
		}
	}()
	return snapshots, errs, nil
}
