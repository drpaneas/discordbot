# Discord Bot

[![Go Report Card](https://goreportcard.com/badge/github.com/drpaneas/discordbot)](https://goreportcard.com/report/github.com/drpaneas/discordbot)
[![golangci-lint](https://github.com/drpaneas/discordbot/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/drpaneas/discordbot/actions/workflows/golangci-lint.yml)
[![Go](https://github.com/drpaneas/discordbot/actions/workflows/test.yml/badge.svg)](https://github.com/drpaneas/discordbot/actions/workflows/test.yml)
[![Every 1 Hour](https://github.com/drpaneas/discordbot/actions/workflows/run.yml/badge.svg)](https://github.com/drpaneas/discordbot/actions/workflows/run.yml)
[![GoDoc](https://godoc.org/github.com/drpaneas/discordbot?status.svg)](https://pkg.go.dev/mod/github.com/drpaneas/discordbot)
[![License](https://img.shields.io/:license-GPL3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

----

Το project αντιπροσωπεύει ένα [Discord] bot το οποίο δημοσιεύει ειδήσεις γύρω από την αστρονομία και την αστροφυσική με στόχο την μετάδοση της γνώσης.

---

## Οδηγίες για build

* Απαιτείται να έχεις εγκατεστημένη τη γλώσσα προγραμματισμού της Go (δες τις [οδηγίες](https://go.dev/doc/install)).
* Στην συνέχεια κατεβάζεις τον κώδικα στον υπολογιστή σου με όποιον τρόπο θες (πχ `git clone https://github.com/drpaneas/discordbot`).
* Και κάνεις build: `go build`

## Πώς να το χρησιμοποιήσετε

### 1ος τρόπος: χωρίς το Discord

Το τρέχεις έτσι όπως θα έτρεχες ένα οποιαδήποτε άλλο πρόγραμμα Go.

```shell
./discordbot # αφού το έχεις κάνει build πρώτα
```

Είτε:

```shell
go run main
```

### 2ος τρόπος: επικοινωνία με το Discord

Για να το τρέξεις, πρέπει ορίσεις τη μεταβλητή `WEBHOOK_DISCORD` που έχεις στο webhook, μέσω της οποίας μπορεί να έχει πρόσβαση στο Discord.
Για να την βρεις, κάνε _"Δεξί Κλικ"_ πάνω στο κανάλι που θες να βάζει τα νέα, (πχ `#Αστρονέα`), επέλεξε `Edit Channel`, πάνε στο `Integrations` και τέλος `Webhooks`.
Εκεί θα μπορείς είτε να φτιάξεις ένα καινούριο Webhook είτε να δεις (_View Webhooks_) τα ήδη υπάρχοντα.

```shell
WEBHOOK_DISCORD="https://discord.com/api/webhooks/βάλε_το_δικό_σου" ./discordbot
```

Στη συνέχεια, θα πρέπει να το βάλεις να τρέχει κάθε 60 λεπτά στο GitHub action.
Για να πάρει την μεταβλητή το GitHub, πάνε στο `Settings > Secrets > Actions` και βάλε ένα `New Repository Secret` με όνομα `WEBHOOK_DISCORD` και value `αυτό που σου δείχνει το discord`.
Τέλος πάτα `Add Secret`.
Για να το κάνει διαθέσιμο στο περιβάλλον που θα τρέχει το job:

```yml
env:
      WEBHOOK_DISCORD: ${{ secrets.WEBHOOK_DISCORD }}
```

Μπορείς να δεις το συγκεκριμένο workflow [εδώ](.github/workflows/run.yml).

## Αδεια χρήσης

To project αποτελεί προϊόν Ελεύθερου και Ανοιχτού Λογισμικού κάτω από την άδεια [GPLv3](https://fsfe.org/activities/gplv3/gplv3.el.html).

## Αποποίηση Ευθύνης

Το _discodbot_ και οι συνεισφέροντες αυτού, **δεν** ενεργούν εξ ονόματός των ειδησιογραφικών websites από τα οποία αντλούνται οι ειδήσεις ή φέρουν ευθύνη για οποιαδήποτε ενδεχόμενη κατάχρηση του συγκεκριμένου project.

Ολα τα άλλα εμπορικά σήματα είναι ιδιοκτησία των αντίστοιχων ιδιοκτητών τους.

[Discord]: https://discord.com/