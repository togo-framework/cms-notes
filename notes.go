// Package notes is the Notes togo plugin: Knowledge base + task board + AES-256-GCM secret vault (PIN-gated, ciphertext never leaves the server), one-time/public share links, and convert-note-to-published-content.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package notes

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("notes", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/notes/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"notes","status":"ok"}`))
		})
		return nil
	})
}
