;; 👋🏽 Hi there! Welcome to the Clojure Track.
;; The online test-runner is powered by babashka and the Small Clojure Interpreter (SCI).
;; Almost all language features are supported, with the exception of low-level constructs
;; like `deftype`, and certain Java classes. For more info, see:
;; https://github.com/babashka/babashka#differences-with-clojure

(ns lucians-luscious-lasagna)

(def expected-time 40
  )

(defn remaining-time
  [actual-time]
  (- expected-time actual-time)
  )

(defn prep-time
  [num-layers]
  (* num-layers 2)
  )

(defn total-time
  [num-layers actual-time]
  (+ actual-time (* num-layers 2))
  )
