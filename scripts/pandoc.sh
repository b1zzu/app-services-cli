files=($(find "./docs/commands" -type f -name '*.md'))
for item in ${files[*]}
do
  filename="`basename ${item%.*}`"   
  adocFile="${item%.*}.adoc"   
  printf "Working on   %s\n" $filename
  pandoc -s $item -f markdown -t asciidoc -o ${DIR}${adocFile}
  echo "Generated adoc. Removing markdown"
  rm -f ${item}
done