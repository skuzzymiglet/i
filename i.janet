(import termbox :as tb)
(import utf8)

(def chars {:top-left "┌"
            :top-right "┐"
            :bottom-left "└"
            :bottom-right "┘"

            :vertical-line "│"
            :horizontal-line "─"

            :vertical-left "┤"
            :vertical-right "├"
            :horizontal-up "┴"
            :horizontal-down "┬"})

(defn char [s] (get s 0))

(defer (tb/shutdown)
  (tb/init)
  (tb/put-string 0 0 (utf8/encode "┌") (tb/color :default) (tb/color :default))
  (tb/present)
  (pp (tb/poll-event)))
