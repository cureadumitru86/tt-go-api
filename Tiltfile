load('ext://nerdctl', 'nerdctl_build')
nerdctl_build(
    ref='go-demo',
    context='.',
    dockerfile='dev.dockerfile',
    live_update=[
        sync('src/cmd', '/app/cmd'),
        sync('src/server', '/app/server'),
        sync('src/go.mod', '/app/go.mod'),
        sync('src/main.go', '/app/main.go')
    ]
)

k8s_yaml(helm("./charts/go-demo", namespace="demo"))

k8s_resource('go-demo',
             # map one or more local ports to ports on your Pod
              port_forwards=['8080:8080'],
             # change whether the resource is started by default
             auto_init=True,
             # control whether the resource automatically updates
             trigger_mode=TRIGGER_MODE_AUTO
)