package cmd

import (
	"fmt"

	"github.com/contextos-hq/untag-cli/internal/claims"
	"github.com/contextos-hq/untag-cli/internal/localconfig"
	"github.com/contextos-hq/untag-cli/internal/ui"
)

// loadLink reads folderPath's .untag (see localconfig.Load), moving a
// claim token written there by an older untag to the user's claims store, and
// tells the user once when it does. It warns if a claim token is, or was,
// exposed to git (untag-wkc).
func loadLink(folderPath string) *localconfig.Config {
	cfg, res := localconfig.Load(folderPath)
	if !res.Migrated && !res.SecretInFile {
		return cfg
	}
	git := localconfig.Git(folderPath)

	if res.Migrated {
		where := "your untag config folder"
		if p, err := claims.Path(); err == nil {
			where = p
		}
		fmt.Printf("\n%s Moved this site's claim token out of %s (a secret: anyone with it can claim the site) to %s.\n", ui.Green("✓"), localconfig.FileName, where)
		if git.Tracked {
			msg := localconfig.FileName + " is committed to git, so earlier commits still contain the claim token: anyone who can read this repository can claim the site."
			if cfg.ClaimURL != "" {
				msg += "\nClaim it now to keep it: " + cfg.ClaimURL
			} else {
				msg += "\nClaim it now to keep it."
			}
			ui.PrintWarning(msg)
		}
		fmt.Println()
	}

	if res.SecretInFile && git.InRepo && !git.Ignored {
		ui.PrintWarning(localconfig.FileName + " contains this site's claim token (a secret: anyone with it can claim the site) and isn't git-ignored.\nAdd " + localconfig.FileName + " to .gitignore so it isn't committed.")
	}
	return cfg
}
