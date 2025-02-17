// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.
/*
 * Parts of this file were generated with makeClass --run. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

package scrapePkg

// EXISTING_CODE
import (
	"net/http"
	"sync"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/colors"
	"github.com/spf13/cobra"
)

// EXISTING_CODE

func RunScrape(cmd *cobra.Command, args []string) error {
	opts := ScrapeFinishParse(args)

	err := opts.ValidateScrape()
	if err != nil {
		return err
	}

	// EXISTING_CODE
	var wg sync.WaitGroup

	wg.Add(1)
	IndexScraper = NewScraper(colors.Yellow, "IndexScraper", opts.Sleep, opts.Globals.LogLevel)
	go RunIndexScraper(opts, wg, hasIndexerFlag(args[0]))

	wg.Add(1)
	MonitorScraper = NewScraper(colors.Purple, "MonitorScraper", opts.Sleep, opts.Globals.LogLevel)
	go RunMonitorScraper(opts, wg, hasMonitorsFlag(args[0]))

	wg.Wait()

	return nil
	// EXISTING_CODE
}

func ServeScrape(w http.ResponseWriter, r *http.Request) bool {
	opts := FromRequest(w, r)

	err := opts.ValidateScrape()
	if err != nil {
		opts.Globals.RespondWithError(w, http.StatusInternalServerError, err)
		return true
	}

	// EXISTING_CODE
	return false
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
