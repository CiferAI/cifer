## ✅ `README.md`

````markdown
# Cifer Blockchain

**Cifer** is a custom blockchain built using [Cosmos SDK](https://docs.cosmos.network) and [Tendermint](https://docs.tendermint.com), scaffolded with the [Ignite CLI](https://ignite.com/cli).  
This chain powers on-chain federated learning metadata and asset minting through the `engrave` module.

---

## 🚀 Get Started

```bash
ignite chain serve
````

This command installs dependencies, builds proto files, initializes configuration, and starts your blockchain in development mode.

---

## ⚙️ Configuration

Edit `config.yml` to customize ports, app settings, or build options.

For more details, refer to the [Ignite CLI documentation](https://docs.ignite.com).

---

## 🌐 Web Frontend Integration

Ignite supports frontend scaffolding using Vue or React:

* Vue:

  ```bash
  ignite scaffold vue
  ```

* React:

  ```bash
  ignite scaffold react
  ```

Learn more in the [Ignite Web Monorepo](https://github.com/ignite/web).

---

## 📌 Usage: Mint Metadata with `engrave` Module (CLI)

You can mint on-chain metadata for federated learning models using the `engrave` module:

```bash
ciferd tx engrave metadata \
  wallet1xyz... \
  "Federated Server v1.2" \
  "ResNet-18" \
  model123 \
  model_weights \
  "N/A" \
  "Cifer.aI" \
  wallet1abc... \
  "Non-commercial" \
  "2025-12-31T00:00:00Z" \
  "2025-07-09T12:00:00Z" \
  "2025-07-09T12:00:00Z" \
  1234 \
  --from mywallet
```

Query the metadata:

```bash
ciferd query engrave metadata model123
```

---

## 📡 Usage: Mint Metadata via REST API (JSON)

If you have REST enabled (via `ignite chain serve`), you can also send metadata as JSON:

### Endpoint

```
POST /cifer/engrave/metadata
```

### Example JSON

```json
{
  "creator_wallet": "wallet1xyz...",
  "generated_via": "Federated Server v1.2",
  "model_used": "ResNet-18",
  "asset_id": "model123",
  "asset_type": "model_weights",
  "prompt_attribution": "N/A",
  "ip_owner_name": "Cifer.aI",
  "ip_owner_wallet": "wallet1abc...",
  "usage_rights": "Non-commercial",
  "license_expiry_date": "2025-12-31T00:00:00Z",
  "created_on": "2025-07-09T12:00:00Z",
  "on_chain_record": "2025-07-09T12:00:00Z",
  "block_height": 1234
}
```

### Example cURL

```bash
curl -X POST http://localhost:1317/cifer/engrave/metadata \
  -H "Content-Type: application/json" \
  -d @metadata.json
```

---

## 📦 Release

To release a new version of your chain:

```bash
git tag v0.1.0
git push origin v0.1.0
```

This creates a draft release on GitHub that you can publish manually.

---

## 📥 Install Binary

To install the latest released version of your blockchain binary:

```bash
curl https://get.ignite.com/CiferAI/cifer@latest! | sudo bash
```

> Make sure `CiferAI/cifer` matches your GitHub repo path.

More info on [Ignite installer](https://github.com/allinbits/starport-installer).

---

## 📚 Learn More

* [Ignite CLI](https://ignite.com/cli)
* [Ignite CLI Docs](https://docs.ignite.com)
* [Cosmos SDK Docs](https://docs.cosmos.network)
* [Ignite Tutorials](https://docs.ignite.com/guide)
* [Ignite Developer Discord](https://discord.gg/ignite)

---

Built with ❤️ by [Cifer.aI](https://github.com/CiferAI)

```