# CURSO DE KUBERNETES

> Este contenido es una toma de notas de un curso desactualizado para el año 2020, el cual puede contener partes erroneas si bien se va chaqueando con la documentación actual.

Para entender mejor el concepto de kubernetes y el principal contenido introductorio, se recomienda ver el video: [![video](https://img.youtube.com/vi/7bA0gTroJjw/0.jpg)](https://www.youtube.com/watch?v=7bA0gTroJjw)

# Índice

* [Introducción](#Introducción)
  * [Comunicación con la API](#comunicación-con-la-api)
  * [Master](#master)
    * [Scheduler](#Scheduler)
    * [Controllers](#Controllers)
    * [etcd](#etcd)
  * [Kubelet](#Kubelet)
  * [Kubeproxy](#Kubeproxy)
  * [Container runtime](#Container-runtime)
* [Pods](#Pods)
  * [Crear](#crear-pod)
  * [Ver](#ver-pods)
  * [Borrar](#borrar-pod)
  * [Info](#Info-del-pod)
  * [Logs](#Ver-logs-del-pod)
  * [Manifiesto](#Manifiesto)
    * [Ejecutar](#Ejecutar-manifiesto)
      * [Ejemplo - Varios contenedores en un pod](#Crear-varios-contenedores-en-un-pod-con-manifest)
    * [Labels](#Labels-en-pods)
  * [Problemas](#Problema-con-los-pods)
* [Objetos de mayor nivel](#OBJETOS-DE-MAYOR-NIVEL)
  * [ReplicaSet](#ReplicaSet)
  * [Deployments](#Deployments)
    * [Actualizar](#Actualizar-deployment)
    * [Rollout](#Rollout) 
    * [Historial y revisiones](#Historial-y-revisiones-de-deploy
) 
    * [CHANGE-CAUSE](#CHANGE-CAUSE) 





  

  

  



# Introducción

## Comunicación con la API

- `kubectl` es el comando para comunicarse con la API de kubernetes, transformando un comando linux en un JSON y pasandoselo a la misma con información y acciones sobre contenedores.

## Master

El master es el "cerebro" de kubernetes, se encarga de orquestar los nodos que pueden verse como MV's. 

Contiene los siguientes componentes: 


### Scheduler

-   Toma las instrucciones de la API y busca los nodos con mejores recursos para desplegar.

- Si no hay nodos disponibles con lo especificado, se queda con el contenedor en estado `P` hasta conseguir los recursos en algun nodo.

### Controllers (kubecontroller)

- `node controller` se encarga de levantar maquinas virtuales/containers que se hayan caido

- `replication controller` se encarga que las replicas configuradas esten siempre corriendo.

- `endpoint controller` servicios y pods

- `service account controller` y `tokens controllers` para autenticación


### etcd

- Base de datos key/value en donde el cluster almacena el estado, los datos, backups, etc.

- Guarda información de despliegues anteriores que facilitan los rollbacks, por ej.


## Kubelet

¿Cómo se comunica el master con los nodos?
Cada nodo (maquina virtual, o no) corre el servicio kubelet.
Este servicio es el responsable de recibir y enviar info de los nodos al master


## Kubeproxy

Es un servicio que se corre tambien en los nodos y maneja todos los temas de red.

## Container runtime

-   Docker



>Se puede ingresar a ver TODOS los recursos de la api con el siguiente comando `kubectl api-resources` así como consultar todos los objetos creados desde `kubectl get all`


# Pods

Es un 'recubrimiento' o 'wrapper' para contenedores en el cual hace que compartan namespaces para comunicarse entre ellos, como: 
-    el UTS (hostname)
-    network (misma IP unica)
-    IPC (pueden ver los procesos entre ellos)

Manteniendo los demas namespaces independientes para cada conenedor (como lo hacen siempre):

-   cgroup (asignar recursos, CPU, ram, etc)
-   PID (procesos)
-   mount (controlar volúmenes, sistema de archivos)   

El pod es un 'wrapper' que permite la interaccion entre ellos.

>Un pod es la unidad mas pequeña de Kubernetes

### Crear Pod

```bash
kubectl run <POD_NAME> --image=<IMAGE_NAME>
```

### Ver pods
Para traer todos
  ```bash
  kubectl get pods
  ```
Para traer uno
  ```bash
  kubectl get pod <POD_NAME>
  ```
Traer toda la info de la generacion del pod en json o yaml (manifiesto)
  ```bash
  kubectl get pod <POD_NAME> -o yaml
  ```

### Borrar pod

```bash
kubectl delete pod <POD_NAME>
```
### Info del pod
```bash
kubectl describe pod <POD_NAME>
```
> Se pueden ver desde acá los datos del pod, como ser la IP, volúmenes, network y los eventos (log)

#### Exponer puerto del pod
```bash
kubectl expose pod <POD_NAME> --type=NodePort --port=80 
```

> `--port` es el puerto que expone el contendor. Se ingresa luego por `localhost:<PORT>`. Consultar el puerto con `kubectl get svc` (rango 30000)

> Según el video del principio para exponer el servicio al exterior es necesario crear un servicio de k8s, al igual que un deployment ya que rara vez se usa el pod solo. Esto es lo que hace el comando anterior, crea un servicio.
  
  

#### Ejecutar comandos o entrar a la consola del pod
```bash
kubectl exec -ti <POD_NAME> -- (sh o comando)
```

#### Ver logs del pod

```bash
Kubectl logs <POD_NAME> 
```

> Flag `-f` para ver los logs en vivo

# Manifiesto

El manifiesto es un archivo .yaml donde se define el recurso de k8s a crear o actualizar.

Ejemplo básico: 

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod #nombre del pod
spec:
  containers:
    - name: web-server #nombre del contenedor
      image: nginx:stable-alpine
      ports:
        - name: web
          containerPort: 82
          protocol: TCP
```

* El `kind` se puede obtener de la ultima columna de los recursos de la API `kubectl api-resources`
* Los specs describen el contenedor que queremos crear

### Ejecutar manifiesto
Para ejecutar este manifiesto debemos usar el comando: 
```bash
kubectl apply -f filename.yaml
```

### Borrar manifiesto
Para borrar lo creado con este manifiesto debemos usar el comando: 
```bash
kubectl delete -f filename.yaml
```


## Crear varios contenedores en un pod con manifest

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: varios-contenedores #nombre del pod
spec:
  containers:
    #Primer container
    - name: server-1 #nombre del contenedor
      image: nginx:stable-alpine
      ports:
        - name: web
          containerPort: 82
          protocol: TCP
    #Segundo container
    - name: server-2 
      image: nginx:stable-alpine
      ...
      ...
      ...
```

#### Comandos en pods con varios contenedores

```bash
kubectl logs varios-contenedores -c server-2
```
```bash
kubectl exec -ti varios-contenedores -c server-2 -- sh
```

## Labels en pods

Es metadata que se le agrega al pod para distinguirlo.
Por ejemplo, para diferenciar pods de produccion, testing, etc.



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: varios-contenedores
  label:
    app: backend
    env: dev
spec:
  containers:
  ...
  ...
```

* Buscar pods por label
```bash
kubectl get pods -l app=backend
```

* Agregar labels a pods

```bash
kubectl label pods <POD_NAME> new-label=value
```

> Los labels son importantes ya que aparte de filtrar las búsquedas permite a los elementos de más alto nivel (deployments, replicasets) administrar los pods

## Problema con los pods

* No pueden recuperarse automaticamente si se 'caen' o son eliminados. Es decir no hacen `self-healing`
* No pueden réplicarse a si mismos
* No pueden actualizarse a si mismos

>Para solucionar este problema se utiizan objetos de k8s de mayor nivel
  
>Los pods deben ser creados por objetos de mayor nivel, ya que puede haber solapamiento de labels con pods creados a mano y los otros recursos los 'adoptan' por tales labels

---

# OBJETOS DE MAYOR NIVEL

## ReplicaSet
Se encarga de crear pods definidos en un template y mantener el numero de replicas activas requeridas del mismo (`self-healing`).

Para que el RS sepa que pods debe manejar, se le pasa los labels de estos y en caso de no existir ninguno con ese label, lo crea.

Aparte el RS va a agregar a la metadata un `owner-reference` que es un identificador que apunta a los pods hacia el. Con esto, los pods tienen como 'dueño' a este RS y no pueden ser tomados por otro.


### Manifiesto
El RS se define de la siguiente manera:

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: rs-test
  labels:
    app: rs-test
spec:
  replicas: 5
  selector:
    matchLabels:
      app: pod-label  #selecciona o crea pods con este label, pueden ser más
  template: # El template es la definición del pod que se quiere crear
    metadata:
      name: rs-test-pod
      labels:
        app: pod-label
    spec:
      containers:
        - name: web-server
          image: nginx
          ports:
            - name: web
              containerPort: 82
              protocol: TCP
```

### Comandos

Se pueden utilizar los comandos `get`, `describe`, `delete`, etc. Solo debemos especificar que el recurso es un RS.

```bash
kubectl get rs
```

### Problemas de los RS

Los RS solo miran los labels de los pods que manejan, NO pueden actualizar los pods para cambiar imágenes (actualizar versión de imagen), configuraciones, etc.

>Se podría forzar la actualización al borrar los pods con configuraciones viejas y dejar que el RS los levante nuevamente, ya con la nueva configuración.

---

## Deployments

Es un objeto de alto nivel que crea un RS, el RS creado al mismo tiempo ejecuta la creación de los pods. 
El deployment se encarga de plasmar los cambios en el estado de los RS, pods u otros deployments. Para esto crea un nuevo RS y se guia por dos parámetros ya definidos (pero modificables):

  * El `max unavailable`: garantiza que sólo un porcentaje de Pods puede eliminarse mientras se están actualizando. Por defecto es el 25%.
  * El `max surge`:  garantiza que sólo un porcentaje de Pods puede crearse por encima del número deseado de Pods. Por defecto es el 25%.

Lo puntos anteriores garantizan que el deployment no elimina los viejos Pods hasta que un número suficiente de nuevos Pods han arrancado, y no crea nuevos Pods hasta que un número suficiente de viejos Pods se han eliminado.

> Los pods nuevos se crean en un RS nuevo escalando al numero de replicas deseado, y el RS viejo escala a cero a medida que se van borrando los pods.

### Manifiesto

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dep-test
  labels:
    app: back
spec:
  replicas: 3
  selector:
    matchLabels:
      app: back
  template:
    metadata:
      labels:
        app: back
    spec:
      containers:
      - name: nginx
        image: nginx:stable-alpine
```

### Actualizar deployment

Si se han realizado cambios en la configuración de los pods o en la versión del deployment, para actualizar los cambios se usa el mismo comando inicial que para levantar recursos: 

```bash
kubectl apply -f deploy-name.yaml
```

Una vez actualizado, se aplicarán las reglas definidas por los parámetros `max unavailable` y `max surge`. 

Para ver que esta pasando en la actualización 

```bash
kubectl rollout status deployment <DEPLOY_NAME>
```
> Para verlo en tiempo real se debe correr justo después del `apply` ya que sino se verá solo el mensaje de rollout exitoso


 O tambíen se pueden ver la sección de eventos desde:

 ```bash
 kubectl describe deployments <DEPLOY_NAME>
 ```
### Rollout

Kubernetes mantiene cierto número RS por deployment y el hecho de guardarlas ayuda a realizar un rollback ante cualquier desperfecto de la versión nueva.

>Se puede cambiar la cantidad de RS almacenadas desde el manifiesto en los `spec` agregando `revisionHistoryLimit` (por defecto es 10)

 ### Historial y revisiones de deploy

 Al realizar una actualización se guardan los RS utilizados anteriormente, esto se puede ver trayendo con `get rs` con los labels que les hallamos asignado.

 Tendremos el RS nuevo con los pods funcionando y los RS anteriores en cero. 

 Para ver el historial o rollouts: 
 
 ```bash
 kubectl rollout history deploy <DEPLOY_NAME>
 ```

### CHANGE-CAUSE
Sirve para identificar que se cambió entre versiones del deploy: 

* Para grabar que comando fue el que realizó el cambio 

  ```bash
  kubectl apply -f dep.yaml --record
  ```

* Para agregar una anotación detallada de los cambios
hay que modificar el manifiesto, agregando en metadata:

  ```yaml
  metadata:
    annotations:
      kubernetes.io/change-cause: "description"
    name: dep-test
    labels:
      app: back
  ```

---

## Services

Un servicio es un objeto que reúne pods según un agrupador (label), exponiendo un punto de salida común (IP) a todos ellos. 

Es especialmente útil para mapear pods de un deployment o replicaset que pueden caerse y regenerarse cambiando las IP's internas de los mismos. El servicio observa los pods con el label o agrupador definido y expone una IP única.

> El servicio va a actuar como balanceador con los pods que esté observando, distribuyendo los requests que ingresan entre estos.

> No importa si los pods pertenecen a distintos deployments o replicasets, el servicio selecciona pods según el label o agrupador definido

> Kubernetes garantiza que el IP del service es inmutable en el tiempo

### Endpoints

Es un objeto que se crea automáticamente al crear un service. Cuando el servicio encuentra un pod que cumpla con el label definido, coloca la IP de este en el endpoint. De esta forma, el servicio sabe a que IP's dirigirse cuando tiene un request.

> Un endpoint de forma simplificada es una lista de las IP's de los pods que cumplen con el label definido.

#### Ver endpoints

```bash
kubectl get endpoints
```

### Definir servicio para un deployment o pod

#### Manifiestio

```yaml
#primero se define un deploy
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dep-test
  labels:
    app: backend-deploy
spec:
  replicas: 3
  selector:
    matchLabels:
      app: backend-pod
  template:
    metadata:
      labels:
        app: backend-pod
    spec:
      containers:
      - name: nginx
        image: nginx:stable-alpine
---
#separando con 3 guiones, se pueden definir varios objetos
#Servicio
apiVersion: v1
kind: Service
metadata:
  name: my-service
  labels:
    app: backend-service
spec:
  selector:
    app: backend-pod #label del pod, no deploy ni rs
  ports:
    - protocol: TCP
      port: 8085 #puerto para exponer
      targetPort: 80 #puerto del contenedor
```


> Con `kubectl get svc my-svc` o `kubectl describe svc my-service` se puede obtener la IP interna asignada y el listado del endpoint

> Si no se especifica el servicio se creará por default como `ClusterIP` y sin salida exterior a kubernetes.

### DNS
Cuando creamos un servicio, este hereda una IP y un DNS, pudiendo usarse indistintamente para hacer llamados al mismo.

Por ejemplo para probar que funcione el servicio descripto por el manifiesto (ClusterIP):

```bash
#crear un pod temporal que se elimina cuando se sale de la terminal
kubectl run --rm -ti test-traffic --image=nginx:stable-alpine -- sh 

#ya adentro del pod
apk add -U curl

#probar IP del servicio (consultar con get o describe)
curl IP:<PORT>

#probar DNS
curl <SVC_NAME>

```

> Los DNS se relacionan con los namespaces, que se ven mas adelante


### Tipos de servicios

* **ClusterIP:** IP virtual que k8s asigna al servicio. Esta es interna al cluster, por lo que no se puede consultar desde afuera del mismo. Es para comunicación interna con componentes del cluster.

* **NodePort:** Es un puerto que se abre a nivel del nodo (máquina virtual, física o conteinerizada) para exponer el servicio externamente al cluster.
  > Si no se especifica puerto, se asigna automáticamente (ver con `get` o `describe`).
  Luego ya se puede ingresar desde `localhost:<PORT>`

* **LoadBalancer:** Crea un balanceador de carga, que expone una IP externa. Al mismo tiempo abre NodePorts en cada nodo. Al acceder al IP se distribuyen las peticiones en los NodePorts. 
  > El servicio se vuelve accesible externamente solo con la funcionalidad de LB que dan los proveedores de Cloud (AWS, GCP, Azure)




![alt text](https://miro.medium.com/max/700/0*-F0LFSm6vKAQcgbq.jpg "Service types graph")




### Asignar tipo de servicio

En el manifiesto, en los specs del servicio, agregar: 

  ```yaml
  #...

  spec:
  type: ClusterIP # or NodePort or LoadBalacer
  selector:
    app: backend-pod
  ports:

  #...
  ```

___

## Namespaces

Son clusters virtuales que se pueden crear dentro del mismo cluster físico, permitiendo que a cada uno se le pueda asignar recursos aislados entre sí. A estos se los llama `namespaces`. 

> Los `namespaces` son una forma de dividir los recursos del clúster entre múltiples usuarios, proyectos o clientes. Define `scopes` para tal motivo.

Los namespaces permiten tambien limitar objetos y recursos de hardware, por ejemplo: 

  * Cantidad de objetos (pods, RS, deploys, etc.)
  * Cantidad de CPU y RAM que consumirán los objetos
  * Limitar acceso de usuarios
  * Limites por defectos


### Ver namespaces

```bash
#abreviatura 'ns'
kubectl get namespaces --show-labels 
```

> Todos los pods que se crean, si no se especifica, lo hacen en el namespace `default`

### Crear namespaces

Se puede crear de dos formas:

  1. 
      ```bash
      kubectl create namespace <NS_NAME>
      ```

  2. 
      ```yaml
      apiVersion: v1
      kind: Namespace
      metadata:
        name: development
        labels:
          name: development
      ```
      Luego:
      ```bash
      kubectl create -f <NS_FILE_NAME>
      #tambien
      kubectl apply -f <NS_FILE_NAME>
      ```

### Describir namespaces

  ```bash
  kubectl describe namespaces <NS_NAME>
  ```

### Ver objetos de namespaces

Para ver los distintos objetos dentro de un namespace: 

```bash
kubectl get all -n <NS_NAME>
# Igual para cada objeto
kubectl get pods -n <NS_NAME>
```

### Crear objetos en namespaces
Para crear objetos en un namespace distinto al elegido como default.

* Desde linea de comandos:

    ```bash
    kubectl run <POD_NAME> --image=<IMAGE_NAME> --namespace <NS_NAME>
    ```
 * O mediante un yaml/json, en `metadata`:

      ```yaml
      kind: Deployment
      metadata:
        name: dep-test
        namespace: <NS_NAME>
        labels:
          app: back
      ```

### Comunicacíon entre distintos namespaces

Como habiamos visto, se puede acceder a un servicio por medio de un DNS que coincide con el nombre del mismo:

```bash
curl <SVC_NAME>
```

Esto solo funciona internamente al namespace, para comunicarse entre los distintos namespaces, el DNS se compone de ciertas partes:

```bash
# Fully qualified domain name or FQDN
curl <SVC_NAME>.<NS_NAME>.svc.cluster.local
```

### Contextos

El contexto es el namespace por defecto. 
Se puede ver que contexto se esta utilizando desde el comando: 

```bash
kubectl config view
```

Si queremos cambiarlo para no agregar `-n <NS_NAME>` a todos los comandos.

#### Crear contexto

```bash
kubectl config set-context <CONTEXT_NAME> \
--namespace=<NS_NAME> \
--cluster=<CLUSTER_NAME> \
--user=<CLUSTER_USER_NAME> 

#en este caso es <CLUSTER_NAME> y <CLUSTER_USER_NAME>  son docker-desktop o minikube
```

#### Usar contexto 

```bash
kubectl config use-context <CONTEXT_NAME>
```
