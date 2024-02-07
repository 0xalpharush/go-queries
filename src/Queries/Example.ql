import go

from Function println, DataFlow::CallNode call
where
  println.hasQualifiedName("fmt", "Printf") and
  call = println.getACall()
select call