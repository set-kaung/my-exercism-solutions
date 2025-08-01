(ns tracks-on-tracks-on-tracks)

(defn new-list
  "Creates an empty list of languages to practice."
  []
  list '()  )

(defn add-language
  "Adds a language to the list."
  [lang-list lang]
  (conj lang-list lang)
  )

(defn first-language
  [lang-list]
  (nth lang-list 0)
  )

(defn remove-language
  "Removes the first language added to the list."
  [lang-list]
  (rest lang-list)
  )

(defn count-languages
  "Returns the total number of languages on the list."
  [lang-list]
  (count lang-list)
  )

(defn learning-list
  "Creates an empty list, adds Clojure and Lisp, removes Lisp, adds
  Java and JavaScript, then finally returns a count of the total number
  of languages."
  []
   (let [languages (list)
        with-clojure-and-lisp (conj languages "Clojure" "Lisp")
        without-lisp (remove #{"Lisp"} with-clojure-and-lisp)
        with-java-and-javascript (conj without-lisp "Java" "JavaScript")
        total-languages (count with-java-and-javascript)]
    total-languages))
  
