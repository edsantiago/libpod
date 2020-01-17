package server

import (
	"net/http"

	"github.com/containers/libpod/pkg/api/handlers"
	"github.com/containers/libpod/pkg/api/handlers/generic"
	"github.com/containers/libpod/pkg/api/handlers/libpod"
	"github.com/gorilla/mux"
)

func (s *APIServer) RegisterContainersHandlers(r *mux.Router) error {
	// swagger:operation POST /containers/create compat createContainer
	// ---
	//   summary: Create a container
	//   tags:
	//    - containers (compat)
	//   produces:
	//   - application/json
	//   parameters:
	//    - in: query
	//      name: name
	//      type: string
	//      description: container name
	//   responses:
	//     '201':
	//         $ref: "#/responses/ContainerCreateResponse"
	//     '400':
	//         "$ref": "#/responses/BadParamError"
	//     '404':
	//         "$ref": "#/responses/NoSuchContainer"
	//     '409':
	//         "$ref": "#/responses/ConflictError"
	//     '500':
	//        "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/create"), APIHandler(s.Context, generic.CreateContainer)).Methods(http.MethodPost)
	// swagger:operation GET /containers/json compat listContainers
	// ---
	//   tags:
	//    - containers (compat)
	// summary: List containers
	// description: Returns a list of containers
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//       Returns a list of containers.
	//        - ancestor=(<image-name>[:<tag>], <image id>, or <image@digest>)
	//        - before=(<container id> or <container name>)
	//        - expose=(<port>[/<proto>]|<startport-endport>/[<proto>])
	//        - exited=<int> containers with exit code of <int>
	//        - health=(starting|healthy|unhealthy|none)
	//        - id=<ID> a container's ID
	//        - is-task=(true|false)
	//        - label=key or label="key=value" of a container label
	//        - name=<name> a container's name
	//        - network=(<network id> or <network name>)
	//        - publish=(<port>[/<proto>]|<startport-endport>/[<proto>])
	//        - since=(<container id> or <container name>)
	//        - status=(created|restarting|running|removing|paused|exited|dead)
	//        - volume=(<volume name> or <mount point destination>)
	// produces:
	// - application/json
	// responses:
	//   '200':
	//        "$ref": "#/responses/DocsListContainer"
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/json"), APIHandler(s.Context, generic.ListContainers)).Methods(http.MethodGet)
	// swagger:operation POST /containers/prune compat pruneContainers
	// ---
	//   tags:
	//    - containers (compat)
	// summary: Delete stopped containers
	// description: Remove containers not in use
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description:  |
	//      Filters to process on the prune list, encoded as JSON (a `map[string][]string`).  Available filters:
	//       - `until=<timestamp>` Prune containers created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	//       - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune containers with (or without, in case `label!=...` is used) the specified labels.
	// produces:
	// - application/json
	// responses:
	//   '200':
	//       "$ref": "#/responses/DocsContainerPruneReport"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/prune"), APIHandler(s.Context, generic.PruneContainers)).Methods(http.MethodPost)
	// swagger:operation DELETE /containers/{nameOrID} compat removeContainer
	// ---
	//   tags:
	//    - containers (compat)
	// summary: Remove a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: force
	//    type: bool
	//    default: false
	//    description: If the container is running, kill it before removing it.
	//  - in: query
	//    name: v
	//    type: bool
	//    default: false
	//    description: Remove the volumes associated with the container.
	//  - in: query
	//    name: link
	//    type: bool
	//    description: not supported
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '409':
	//       "$ref": "#/responses/ConflictError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}"), APIHandler(s.Context, generic.RemoveContainer)).Methods(http.MethodDelete)
	// swagger:operation GET /containers/{nameOrID}/json compat getContainer
	// ---
	//   tags:
	//    - containers (compat)
	// summary: Inspect container
	// description: Return low-level information about a container.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or id of the container
	//  - in: query
	//    name: size
	//    type: bool
	//    default: false
	//    description: include the size of the container
	// produces:
	// - application/json
	// responses:
	//   '200':
	//      "$ref": "#/responses/DocsContainerInspectResponse"
	//   '404':
	//     "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/json"), APIHandler(s.Context, generic.GetContainer)).Methods(http.MethodGet)
	// swagger:operation POST /containers/{nameOrID}/kill compat killContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Kill container
	// description: Signal to send to the container as an integer or string (e.g. SIGINT)
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: signal
	//    type: int
	//    description: signal to be sent to container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '409':
	//       "$ref": "#/responses/ConflictError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/kill"), APIHandler(s.Context, generic.KillContainer)).Methods(http.MethodPost)
	// swagger:operation GET /containers/{nameOrID}/logs compat logsFromContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Get container logs
	// description: Get stdout and stderr logs from a container.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: follow
	//    type: bool
	//    description: Keep connection after returning logs.
	//  - in: query
	//    name: stdout
	//    type: bool
	//    description: not supported
	//  - in: query
	//    name: stderr
	//    type: bool
	//    description: not supported?
	//  - in: query
	//    name: since
	//    type:  string
	//    description: Only return logs since this time, as a UNIX timestamp
	//  - in: query
	//    name: until
	//    type:  string
	//    description: Only return logs before this time, as a UNIX timestamp
	//  - in: query
	//    name: timestamps
	//    type: bool
	//    default: false
	//    description: Add timestamps to every log line
	//  - in: query
	//    name: tail
	//    type: string
	//    description: Only return this number of log lines from the end of the logs
	//    default: all
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description:  logs returned as a stream in response body.
	//   '404':
	//      "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/logs"), APIHandler(s.Context, generic.LogsFromContainer)).Methods(http.MethodGet)
	// swagger:operation POST /containers/{nameOrID}/pause compat pauseContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Pause container
	// description: Use the cgroups freezer to suspend all processes in a container.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/pause"), APIHandler(s.Context, handlers.PauseContainer)).Methods(http.MethodPost)
	r.HandleFunc(VersionedPath("/containers/{name:..*}/rename"), APIHandler(s.Context, handlers.UnsupportedHandler)).Methods(http.MethodPost)
	// swagger:operation POST /containers/{nameOrID}/restart compat restartContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Restart container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: t
	//    type: int
	//    description: timeout before sending kill signal to container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/restart"), APIHandler(s.Context, handlers.RestartContainer)).Methods(http.MethodPost)
	// swagger:operation POST /containers/{nameOrID}/start compat startContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Start a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: detachKeys
	//    type: string
	//    description: needs description
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '304':
	//       "$ref": "#/responses/ContainerAlreadyStartedError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/start"), APIHandler(s.Context, handlers.StartContainer)).Methods(http.MethodPost)
	// swagger:operation GET /containers/{nameOrID}/stats compat statsContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Get stats for a container
	// description: This returns a live stream of a container’s resource usage statistics.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: stream
	//    type: bool
	//    default: true
	//    description: Stream the output
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/stats"), APIHandler(s.Context, generic.StatsContainer)).Methods(http.MethodGet)
	// swagger:operation POST /containers/{nameOrID}/stop compat stopContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Stop a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: t
	//    type: int
	//    description: number of seconds to wait before killing container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '304':
	//       "$ref": "#/responses/ContainerAlreadyStoppedError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/stop"), APIHandler(s.Context, handlers.StopContainer)).Methods(http.MethodPost)
	// swagger:operation GET /containers/{nameOrID}/top compat topContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: List processes running inside a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: ps_args
	//    type: string
	//    description: arguments to pass to ps such as aux. Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used.
	// produces:
	// - application/json
	// responses:
	//   '200':
	//       "ref": "#/responses/DockerTopResponse"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/top"), APIHandler(s.Context, handlers.TopContainer)).Methods(http.MethodGet)
	// swagger:operation POST /containers/{nameOrID}/unpause compat unpauseContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Unpause container
	// description: Resume a paused container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/unpause"), APIHandler(s.Context, handlers.UnpauseContainer)).Methods(http.MethodPost)
	// swagger:operation POST /containers/{nameOrID}/wait compat waitContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Wait on a container to exit
	// description: Block until a container stops, then returns the exit code.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: condition
	//    type: string
	//    description: Wait until the container reaches the given condition
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     $ref: "#/responses/ContainerWaitResponse"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/wait"), APIHandler(s.Context, generic.WaitContainer)).Methods(http.MethodPost)
	// swagger:operation POST /containers/{nameOrID}/attach compat attachContainer
	// ---
	// tags:
	//   - containers (compat)
	// summary: Attach to a container
	// description: Hijacks the connection to forward the container's standard streams to the client.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: detachKeys
	//    required: false
	//    type: string
	//    description: keys to use for detaching from the container
	//  - in: query
	//    name: logs
	//    required: false
	//    type: bool
	//    description: Not yet supported
	//  - in: query
	//    name: stream
	//    required: false
	//    type: bool
	//    default: true
	//    description: If passed, must be set to true; stream=false is not yet supported
	//  - in: query
	//    name: stdout
	//    required: false
	//    type: bool
	//    description: Attach to container STDOUT
	//  - in: query
	//    name: stderr
	//    required: false
	//    type: bool
	//    description: Attach to container STDERR
	//  - in: query
	//    name: stdin
	//    required: false
	//    type: bool
	//    description: Attach to container STDIN
	// produces:
	// - application/json
	// responses:
	//   '101':
	//     description: No error, connection has been hijacked for transporting streams.
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//       "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/attach"), APIHandler(s.Context, handlers.AttachContainer)).Methods(http.MethodPost)
	// swagger:operation POST /containers/{nameOrID}/resize compat resizeContainer
	// ---
	// tags:
	//  - containers (compat)
	// summary: Resize a container's TTY
	// description: Resize the terminal attached to a container (for use with Attach).
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: h
	//    type: int
	//    required: false
	//    description: Height to set for the terminal, in characters
	//  - in: query
	//    name: w
	//    type: int
	//    required: false
	//    description: Width to set for the terminal, in characters
	// produces:
	// - application/json
	// responses:
	//   '200':
	//        description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//       "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/containers/{name:..*}/resize"), APIHandler(s.Context, handlers.ResizeContainer)).Methods(http.MethodPost)

	/*
		libpod endpoints
	*/

	// swagger:operation POST /libpod/containers/create libpod libpodCreateContainer
	r.HandleFunc(VersionedPath("/libpod/containers/create"), APIHandler(s.Context, libpod.CreateContainer)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/containers/json libpod libpodListContainers
	// ---
	// tags:
	//  - containers
	// summary: List containers
	// description: Returns a list of containers
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     schema:
	//       "$ref": "#/responses/LibpodListContainersResponse"
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/json"), APIHandler(s.Context, libpod.ListContainers)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/containers/prune libpod libpodPruneContainers
	// ---
	// tags:
	//  - containers
	// summary: Prune unused containers
	// description: Remove stopped and exited containers
	// parameters:
	//  - in: query
	//    name: force
	//    type: bool
	//    description: something
	//  - in: query
	//    name: filters
	//    type: string
	//    description:  |
	//      Filters to process on the prune list, encoded as JSON (a `map[string][]string`).  Available filters:
	//       - `until=<timestamp>` Prune containers created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	//       - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune containers with (or without, in case `label!=...` is used) the specified labels.
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: to be determined
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/prune"), APIHandler(s.Context, libpod.PruneContainers)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/containers/showmounted libpod libpodShowMountedContainers
	// ---
	// tags:
	//  - containers
	// summary: Show mounted containers
	// description: Lists all mounted containers mount points
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: mounted containers
	//     schema:
	//      type: object
	//      additionalProperties:
	//       type: string
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/showmounted"), APIHandler(s.Context, libpod.ShowMountedContainers)).Methods(http.MethodGet)
	// swagger:operation DELETE /libpod/containers/{nameOrID} libpod libpodRemoveContainer
	// ---
	// tags:
	//  - containers
	// summary: Delete container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: force
	//    type: bool
	//    description: need something
	//  - in: query
	//    name: v
	//    type: bool
	//    description: delete volumes
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '409':
	//       "$ref": "#/responses/ConflictError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}"), APIHandler(s.Context, libpod.RemoveContainer)).Methods(http.MethodDelete)
	// swagger:operation GET /libpod/containers/{nameOrID}/json libpod libpodGetContainer
	// ---
	// tags:
	//  - containers
	// summary: Inspect container
	// description: Return low-level information about a container.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: size
	//    type: bool
	//    description: display filesystem usage
	// produces:
	// - application/json
	// responses:
	//   '200':
	//       "$ref": "#/responses/LibpodInspectContainerResponse"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/json"), APIHandler(s.Context, libpod.GetContainer)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/containers/{nameOrID}/kill libpod libpodKillContainer
	// ---
	// tags:
	//  - containers
	// summary: Kill container
	// description: send a signal to a container, defaults to killing the container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: signal
	//    type: int
	//    default: 15
	//    description: signal to be sent to container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '409':
	//       "$ref": "#/responses/ConflictError"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/kill"), APIHandler(s.Context, libpod.KillContainer)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/containers/{nameOrID}/mount libpod libpodMountContainer
	// ---
	// tags:
	//  - containers
	// summary: Mount a container
	// description: Mount a container to the filesystem
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: mounted container
	//     schema:
	//      description: id
	//      type: string
	//      example: 3c784de79b791b4ebd3ac55e511f97fedc042328499554937a3f8bfd9c1a2cb8
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/mount"), APIHandler(s.Context, libpod.MountContainer)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/containers/{nameOrID}/logs libpod libpodLogsFromContainer
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/logs"), APIHandler(s.Context, libpod.LogsFromContainer)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/containers/{nameOrID}/pause libpod pauseContainer
	// ---
	// tags:
	//  - containers
	// summary: Pause a container
	// description: Use the cgroups freezer to suspend all processes in a container.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/pause"), APIHandler(s.Context, handlers.PauseContainer)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/containers/{nameOrID}/restart libpod restartContainer
	// ---
	// tags:
	//  - containers
	// summary: Restart a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: t
	//    type: int
	//    description: timeout before sending kill signal to container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/restart"), APIHandler(s.Context, handlers.RestartContainer)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/containers/{nameOrID}/start libpod startContainer
	// ---
	// tags:
	//  - containers
	// summary: Start a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: detachKeys
	//    type: string
	//    description: needs description
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '304':
	//       "$ref": "#/responses/ContainerAlreadyStartedError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/start"), APIHandler(s.Context, handlers.StartContainer)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/containers/{nameOrID}/stats libpod statsContainer
	// ---
	// tags:
	//  - containers
	// summary: Get stats for a container
	// description: This returns a live stream of a container’s resource usage statistics.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: stream
	//    type: bool
	//    default: true
	//    description: Stream the output
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/stats"), APIHandler(s.Context, generic.StatsContainer)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/containers/{nameOrID}/top libpod topContainer
	//
	// List processes running inside a container. Note
	//
	// ---
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: stream
	//    type: bool
	//    default: true
	//    description: Stream the output
	//    name: ps_args
	//    type: string
	//    description: arguments to pass to ps such as aux. Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used.
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: no error
	//       "ref": "#/responses/DockerTopResponse"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/top"), APIHandler(s.Context, handlers.TopContainer)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/containers/{nameOrID}/unpause libpod unpauseContainer
	// ---
	// tags:
	//  - containers
	// summary: Unpause Container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/unpause"), APIHandler(s.Context, handlers.UnpauseContainer)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/containers/{nameOrID}/wait libpod libpodWaitContainer
	// ---
	// tags:
	//  - containers
	// summary: Wait on a container to exit
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: condition
	//    type: string
	//    description: Wait until the container reaches the given condition
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/wait"), APIHandler(s.Context, libpod.WaitContainer)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/containers/{nameOrID}/exists libpod libpodContainerExists
	// ---
	// tags:
	//  - containers
	// summary: Check if container exists
	// description: Quick way to determine if a container exists by name or ID
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: container exists
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/exists"), APIHandler(s.Context, libpod.ContainerExists)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/containers/{nameOrID}/stop libpod stopContainer
	// ---
	// tags:
	//  - containers
	// summary: Stop a container
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: t
	//    type: int
	//    description: number of seconds to wait before killing container
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '304':
	//       "$ref": "#/responses/ContainerAlreadyStoppedError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/stop"), APIHandler(s.Context, handlers.StopContainer)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/containers/{nameOrID}/attach libpod attachContainer
	// ---
	// tags:
	//   - containers
	// summary: Attach to a container
	// description: Hijacks the connection to forward the container's standard streams to the client.
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: detachKeys
	//    required: false
	//    type: string
	//    description: keys to use for detaching from the container
	//  - in: query
	//    name: logs
	//    required: false
	//    type: bool
	//    description: Not yet supported
	//  - in: query
	//    name: stream
	//    required: false
	//    type: bool
	//    default: true
	//    description: If passed, must be set to true; stream=false is not yet supported
	//  - in: query
	//    name: stdout
	//    required: false
	//    type: bool
	//    description: Attach to container STDOUT
	//  - in: query
	//    name: stderr
	//    required: false
	//    type: bool
	//    description: Attach to container STDERR
	//  - in: query
	//    name: stdin
	//    required: false
	//    type: bool
	//    description: Attach to container STDIN
	// produces:
	// - application/json
	// responses:
	//   '101':
	//     description: No error, connection has been hijacked for transporting streams.
	//   '400':
	//       "$ref": "#/responses/BadParamError"
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//       "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/attach"), APIHandler(s.Context, handlers.AttachContainer)).Methods(http.MethodPost)
	// swagger:operation POST /libpod/containers/{nameOrID}/resize libpod resizeContainer
	// ---
	// tags:
	//  - containers
	// summary: Resize a container's TTY
	// description: Resize the terminal attached to a container (for use with Attach).
	// parameters:
	//  - in: path
	//    name: nameOrID
	//    required: true
	//    description: the name or ID of the container
	//  - in: query
	//    name: h
	//    type: int
	//    required: false
	//    description: Height to set for the terminal, in characters
	//  - in: query
	//    name: w
	//    type: int
	//    required: false
	//    description: Width to set for the terminal, in characters
	// produces:
	// - application/json
	// responses:
	//   '200':
	//        description: no error
	//   '404':
	//       "$ref": "#/responses/NoSuchContainer"
	//   '500':
	//       "$ref": "#/responses/InternalError"
	r.HandleFunc(VersionedPath("/libpod/containers/{name:..*}/resize"), APIHandler(s.Context, handlers.ResizeContainer)).Methods(http.MethodPost)
	return nil
}
