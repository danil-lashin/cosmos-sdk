<!--
order: 5
-->

# Governance Proposals

In uncommon situations, a highly valued client may become frozen due to uncontrollable 
circumstances. A highly valued client, might have hundreds of channels being actively used.
Some of those channels might have a significant amount of locked tokens used for ICS 20.

If the one third of the validator set the chain the client represents decides to collude,
they can sign off on two valid but conflicting headers each signed by the other one third
of the honest validator sets. The light client can now be updated with two valid, but conflicting
headers at the same height. The light client cannot know which header is trustworthy and therefore
evidence of such misbehaviour is likely to be submitted resulting in a frozen light client. 

Frozen light clients cannot be updated under any circumstance except via a governance proposal.
Since validators can arbitarily agree to make state transitions that defy the written code, a 
governance proposal has been added to ease the complexity of unfreezing or updating clients
which have become "stuck". Unfreezing clients, re-enables all of the channels built upon that
client. This may result in recovery of otherwise lost funds. 

Tendermint light clients may also become expired if the trusting period has passed since their 
last update. If a chain undergoes an unplanned upgrade, there may be no commitment to that upgrade
signed by the validator set before the chain-id changes. In this situation, the validator set of
the last valid update for the light client is never expected to produce another valid header since 
the chain-id has changed, which will ultimately lead the on-chain light client to become expired.  

In the case that a highly valued light client is frozen, expired, or rendered non-updateable, a
governance proposal may be submitted to update this client, known as the subject client. The 
proposal includes the client identifier for the subject, the client identifier for a substitute
client, and an initial height to reference the substitute client from. Light client implementations
may implement custom updating logic, but in most cases, the subject will be updated with information
from the substitute client, if the proposal passes. The substitute client is used as a "stand in"
while the subject is on trial. It is best practice to create a substitute client *after* the subject
has become frozen to avoid the substitute from also becoming frozen. An active substitute client
allows headers to be submitted during the voting period to prevent accidental expiry once the proposal
passes. 
