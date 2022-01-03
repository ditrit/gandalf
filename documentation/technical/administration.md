# Gandalf - Administration

## Super Administrators

The super-administrators are the highest-level Gandalf users. They are the only ones who can handle the tenants. Super Administrator accounts can only be local accounts of the Gandalf solution. Unlike user accounts defined within a tenant (with the notable exception of tenant administrators), they cannot be synchronized with a corporate year where to use identity disclosure.

## Roles and functions

Super Administrators can:

+ add/remove tenants

+ performing the technical administration gestures at the cluster level, i.e.:

    + adding/removing new members to the cluster

    + shutting down/restarting/upgrading the version of a Gandalf cluster as a whole

    + manage the configuration of the cluster

    + create/remove other super administrators

Super Administrators cannot:

+ action, whatever it may be, within a particular tenant. This role is vested in administrators and tenant users.

## Contingency Account

A Gandalf system always has at least two super-administrator accounts, so that there is necessarily a backup account in case of unavailability of the natural person or password associated with a super-administrator account. It is therefore impossible to delete a Super Administrator account if there are not at least 3 on the system.

## Account Management

A first super-administrator account is automatically created when starting the first cluster member of a Gandalf system

This account is named Administrator by default.

The secret to be used during the first connection as a super-administrator is displayed on the standard output when initializing the first member of the cluster.

During its first connection, a super administrator must, as a first action, change its password (the hardness of the password is controlled).

Once their password is changed, the Super Administrator can manage Super Administrator accounts. They can:

+ Change your own password

+ Create a new super administrator account by assigning it the identifier of its choice

+ Delete a Super Administrator account if there are more than 2 existing Super Administrator accounts

Creating a super-administrator account generates a temporary password. As with the first Super Administrator, any new Super Administrator must change their password on their first login.

## Adding/removing cluster members

A super administrator can declare the addition of "a member to the cluster. Multiple cluster members can be declared without having been created/started yet. It is thus possible to dynamically manage the adaptation to the load of the cluster by starting only some of the declared members.

The declaration of a cluster member generates a secret dedicated to that cluster member. This secret must be inserted in the configuration of the new member of the cluster.

If the secret specified in the configuration of a cluster member does not correspond to any cluster member declared by a super administrator or if it has already been used, its initialization fails.

A super administrator can delete a cluster member. This deletion consists in deleting its declaration (its secret is no longer recognized by the other members of the cluster). As soon as a member of the cluster is deleted, it is automatically stopped and no message can be delivered or accepted from it.

It is not possible to delete the last member of a cluster.


## Tenant Management

Only superadministrators can create and configure tenants.

Tenant management is only possible if at least two super administrator accounts exist.

Superadministrators may only delete a tenant provided that it no longer contains any aggregator.

A local tenant administrator account is created automatically when a tenant is created. Its identifier, within the tenant is Administrator and a temporary password is generated for this account when the tenant is created. This password must be changed during the first connection (the hardness of the password is checked).

A aggregater is also created automatically when creating a tenant.

Within the tenant, a default root application domain is created. It corresponds to the tenant itself, that is to say the perimeter of all the resources that will constitute the tenant.

## Rebooting/reloading a Gandalf cluster

A super administrator can cause the cluster to restart. This will cause each member of the cluster to restart one behind the other. When it is restarted, the member of the cluster updates itself if necessary and reloads its configuration. This mechanism makes it easy to deploy Gandalf cluster updates without interruption of service.

The update of the other Gandalf components (aggregators and connectors) is done at the choice of the maintainers or, depending on their configuration, auomatiquement in reaction to the update of the cluster automatically. The same principle of restarting in sequence the different instances of a component (aggregator or connector) ensures transparent and uninterrupted version updates for system users.

Updates can be scheduled on a one-time or recurring basis automatically (auto update mechanism).

## Cluster Level Configuration

Super administrators can change the configuration of the cluster.

These include the following:

+ Name of the cluster

+ Secondary Database Configuration [Not Existing for MVP]

+ The base URL of the Gandalf repository (by default the Gandalf git)

## The administrators of holding

Tenant administrators are users with the highest level of rights within a tenant of a Gandalf system. They are the only ones who can manage the aggregators of the tenant.

With the exception of Super Administrators, all users of a Gandalf system, including Tenant Administrators, exist within a single Tenant.

The first tenant administrator account created for a (tenant)[#Gestion-des-tenants] is a local account whose identifier is "Administrator". This particular account is created at the same time as it itself. Its password must be changed at the first connection.


## Roles and functions Tenant administrators

Tenant Directors may:

+ Create/remove other tenant administrators

+ Manage the tenant configuration

+ Manage tenant users

+ Declare and manage the aggregators of the tenant

+ Restart/update the tenant’s Gandalf components as a whole

## Management of tenant administrators

A tenant administrator account allows the user to manage the tenant administrator accounts of the tenant. This management follows the same principles as super administrator accounts:

+ Tenant administrator accounts are local accounts

+ At least one additional backup administrator account must be created before other types of resources can be managed

+ A tenant administrator may delete a tenant administrator account only if at least 2 other tenant administrator accounts exist

Unlike other user accounts within the same tenant, the tenant administrator accounts cannot be synchronized with a corporate year to use the identity disclaimer.

## Manage the tenant configuration

Tenant administrators can manage the tenant configuration.

These include the following:

+ Name of tenant

+ Configuration of the authentication chain (local, ldap/ad, delegation)

+ The url of the Gandalf repository for the tenant (by default the one defined at the cluster level)

## Management of tenant users

Tenant administrators can manage (create, edit, delete) tenant users.

The method of authentication and the definition of the accreditation documents used to connect by users depends on the type of authentication used (cf. authentication chain).

By default, the users created by the tenant administrators are root administrators for the tenant.

The role of tenant root administrator allows a user to:

+ Define/modify/delete application domains (which are sub-domains of the tenant’s root domain)

+ Define/modify/delete users with rights to the tenant’s application domains

+ If the necessary rights have been granted, a domain user can manage the Gandalf connectors.

However, domain users cannot create or delete aggregators under any circumstances. only holding administrators can.


## Management of aggressors

An aggregator may be part of only one unit and is characterized within its unit by its logical name;

An aggregator consists of a set of aggregator instances. Aggregator instances have the same role. Using multiple instances for an aggregator simply ensures the resilience, high availability and load bearing of the aggregator.

Each aggregator instance has its own secret in its configuration, which must have been declared in advance.

Several aggregator intances can be declared without having yet been created/started. It is thus possible to dynamically manage the adaptation to the load of the cluster by starting only some of the declared instances of the aggregator.

It is the tenant administrators who declare the addition of "an aggregator instance. This declaration generates the secrecy of the instance. This secret is inserted in the instance configuration.

If the secret specified in the configuration of an aggregator instance does not match any of the instances declared for that aggregator or if it has already been used by another instakce, its startup fails.

A super tenant administrator can delete an instance of agragetaur. This deletion consists in deleting his declaration (his secret is no longer recognized by the Gandalf system). Once an aggregator instance is deleted, it is automatically stopped and no message can be delivered or accepted from it.

It is not possible to delete the last instance of an aggregator that no connector is declared for that aggregator.

## Rebooting/reloading a unit

A tenant administrator can restart all Gandalf aggregators of its tenant. This will cause each instance of each aggregator to restart behind each other. When it is restarted, the cluster instance updates if necessary and reloads its configuration. This mechanism makes it easy to deploy Gandalf aggregator updates without interruption of service.

The updating of the aggregator once done, the updating of the connectors can be done according to the same principle (TODO: digging the automatic scenarios, via the handlers of connexters, etc.)

Tenant administrators can schedule automatic updates (auto update mechanism) on an ad hoc or recurring basis.

## The role of administrator within a tenant

They make it possible to manage domains, users, application contexts, the definition of authorisations and connectors. Depending on the roles assigned, an administrator may or may not manage a type of resource. The default administrative roles in a Gandalf system are:

+ useradmin: on its access perimeter (which is a sub-domain of the tenant), it can:

    + create and manage users,

    + create and manage sub-domains within its scope of access

    + create and manage roles

    + assign administrative and access rights to the domains within its scope of access to users.

+ connectoradmin: it has the rights to declare and manage (including deletion) connectors on a dooned application domain. Declaring, initializing, updating a connector follows exactly the same principles as described for cluster members and aggregator instances.

+ rootadmin: this role is a syntactic sugar to designate a user who has the following roles on the tenant:

    + useradmin with access perimeter holding it itself;

    + connectoradmin on the whole of the holder

it may delegate all or part of its rights to users of the tenant