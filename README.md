# tsproxy

A dead-simple HTTP proxy that creates a [Tailscale](https://tailscale.com/) [private service](https://tailscale.com/blog/tsnet-virtual-private-services/).

## Usage

To create a service named `app` that will proxy a service running on the port `8814`:

```
tsproxy --origin=http://localhost:8814 --hostname app
```

You can access this URL on your browser `http://app`.

Note: [MagicDNS](https://tailscale.com/kb/1081/magicdns/) must be enabled.
